package main

import "net/http"

func (app *AppConfig) SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}
