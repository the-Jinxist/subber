package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

var pathToTemplate = "./cmd/web/templates"

type TemplateData struct {
	StringMap     map[string]string
	IntMap        map[string]int64
	FloatMap      map[string]float64
	Data          map[string]any
	Flash         string
	Warning       string
	Error         string
	Authenticated bool
	Now           time.Time
	//User data.User
}

func (app *AppConfig) render(w http.ResponseWriter, r *http.Request, templatName string, td *TemplateData) {
	partials := []string{
		fmt.Sprintf("%s/base.layout.gohtml", pathToTemplate),
		fmt.Sprintf("%s/alerts.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/footer.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/header.partial.gohtml", pathToTemplate),
		fmt.Sprintf("%s/navbar.partial.gohtml", pathToTemplate),

		// fmt.Sprintf("%s/home.page.gohtml", pathToTemplate),
		// fmt.Sprintf("%s/login.page.gohtml", pathToTemplate),
		// fmt.Sprintf("%s/register.page.gohtml", pathToTemplate),

		// fmt.Sprintf("%s/mail.html.gohtml", pathToTemplate),
		// fmt.Sprintf("%s/mail.plain.gohtml", pathToTemplate),
	}

	var templateSlice []string
	templateSlice = append(templateSlice, fmt.Sprintf("%s/%s", pathToTemplate, templatName))

	for _, v := range partials {
		templateSlice = append(templateSlice, v)
	}

	if td == nil {
		td = &TemplateData{}
	}

	tmpl, err := template.ParseFiles(templateSlice...)
	if err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, app.addDefaultData(td, r)); err != nil {
		app.ErrorLog.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *AppConfig) addDefaultData(td *TemplateData, r *http.Request) *TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.Error = app.Session.PopString(r.Context(), "error")

	if app.isAuthenticated(r) {
		td.Authenticated = true

		//TODO: add more info if user is authenticated
	}
	td.Now = time.Now()

	return td

}

func (app *AppConfig) isAuthenticated(r *http.Request) bool {
	return app.Session.Exists(r.Context(), "userID")
}
