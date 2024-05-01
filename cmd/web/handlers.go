package main

import "net/http"

func (app *AppConfig) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *AppConfig) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *AppConfig) PostLoginPage(w http.ResponseWriter, r *http.Request) {

}

func (app *AppConfig) Logout(w http.ResponseWriter, r *http.Request) {

}

func (app *AppConfig) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *AppConfig) PostRegisterPage(w http.ResponseWriter, r *http.Request) {

}

func (app *AppConfig) ActivateAccount(w http.ResponseWriter, r *http.Request) {

}
