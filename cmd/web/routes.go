package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *AppConfig) Routes() http.Handler {
	//create router
	r := chi.NewRouter()

	//setu middleware
	r.Use(middleware.Recoverer)
	r.Use(app.SessionLoad)

	//define app route
	r.Get("/", app.HomePage)
	r.Get("/login", app.LoginPage)
	r.Post("/login", app.PostLoginPage)
	r.Get("/logout", app.Logout)

	r.Get("/register", app.RegisterPage)
	r.Post("/register", app.PostRegisterPage)

	r.Get("/activate-account", app.ActivateAccount)

	return r
}
