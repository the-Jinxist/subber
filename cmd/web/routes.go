package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *AppConfig) Routes() http.Handler {
	// create router
	mux := chi.NewRouter()

	// set up middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.SessionLoad)

	// define application routes
	mux.Get("/", app.HomePage)

	mux.Get("/login", app.LoginPage)
	mux.Post("/login", app.PostLoginPage)
	mux.Get("/logout", app.Logout)
	mux.Get("/register", app.RegisterPage)
	mux.Post("/register", app.PostRegisterPage)
	mux.Get("/activate", app.ActivateAccount)

	authedRouter := app.authRouter(mux)
	authedRouter.Get("/plans", app.ChooseSubscription)
	authedRouter.Get("/subscribe", app.SubscribeToPlan)

	return mux
}

func (app *AppConfig) authRouter(mux *chi.Mux) chi.Router {
	return mux.With(app.Auth)
}
