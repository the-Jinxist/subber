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

	//define app route
	r.Get("/", app.HomePage)
	return r
}
