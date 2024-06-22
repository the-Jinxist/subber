package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/the-Jinxist/subber/data"
)

func (app *AppConfig) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *AppConfig) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *AppConfig) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	// parse form post
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// get email and password from form post
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {

		msg := Message{
			To:      email,
			Subject: "Failed Login Attempt",
			Data:    "Invalid login attempt!",
		}

		app.sendEmail(msg)

		app.Session.Put(r.Context(), "error", "Invalid credentials.")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// okay, so log user in
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)

	app.Session.Put(r.Context(), "flash", "Successful login!")

	// redirect the user
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *AppConfig) Logout(w http.ResponseWriter, r *http.Request) {
	app.Session.Destroy(r.Context())
	app.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *AppConfig) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *AppConfig) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	//TODO - validate data

	// create user
	u := data.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
	}

	if _, err = u.Insert(u); err != nil {
		app.Session.Put(r.Context(), "error", "Unable to create user")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
	}

	//sending activation email
	url := fmt.Sprintf("http://localhost/activate?email=%s", u.Email)
	signedUrl := GenerateTokenFromString(url)
	app.InfoLog.Println(signedUrl)

	msg := Message{
		To:       u.Email,
		Subject:  "Activate your account",
		Template: "confirmation-email",
		Data:     template.HTML(signedUrl),
	}

	app.sendEmail(msg)

	app.Session.Put(r.Context(), "flash", "Confirmation email sent. Check your email")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

}

func (app *AppConfig) ActivateAccount(w http.ResponseWriter, r *http.Request) {

	// valiadate url
	uri := r.RequestURI // this returns the `?param=value` part of the url
	testUrl := fmt.Sprintf("http://localhost%s", uri)

	okay := VerifyToken(testUrl)
	if !okay {
		app.Session.Put(r.Context(), "error", "Invalid token")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}

	// activate the user

	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(r.Context(), "error", "No user found")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	u.Active = 1
	if err = u.Update(); err != nil {
		app.Session.Put(r.Context(), "error", "Unable to update user")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "flash", "Account Activated. You can now log in")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	// generate an invoice

	// send an email with attachments

	//subscribe the user to an account

}

func (app *AppConfig) ChooseSubscription(w http.ResponseWriter, r *http.Request) {
	if !app.Session.Exists(r.Context(), "userID") {
		app.Session.Put(r.Context(), "warning", "You must login to see this page!")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	dataMap := make(map[string]any)
	dataMap["plans"] = plans

	app.render(w, r, "plans.page.gohtml", &TemplateData{
		Data: dataMap,
	})
}

func (app *AppConfig) SubscribeToPlan(w http.ResponseWriter, r *http.Request) {

	// get the id of the plan that is
	id := r.URL.Query().Get("id")

	planID, _ := strconv.Atoi(id)
	plan, err := app.Models.Plan.GetOne(planID)

	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to find plan")
		http.Redirect(w, r, "/plans", http.StatusSeeOther)
		app.ErrorLog.Println(err)
		return
	}

	// get user form the session
	user, ok := app.Session.Get(r.Context(), "user").(data.User)
	if !ok {
		app.Session.Put(r.Context(), "error", "Login first")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		app.ErrorLog.Println(err)
		return
	}

	app.Wait.Add(1)

	go func() {
		defer app.Wait.Done()

		invoice, err := app.getInvoice(user, plan)
		if err != nil {
			// send this to a channel
			app.ErrorChan <- err
		}

		msg := Message{
			To:       user.Email,
			Subject:  "Your invoice",
			Data:     invoice,
			Template: "invoice",
		}

		app.sendEmail(msg)
	}()

}

func (app *AppConfig) getInvoice(_ data.User, plan *data.Plan) (string, error) {
	return plan.PlanAmountFormatted, nil
}
