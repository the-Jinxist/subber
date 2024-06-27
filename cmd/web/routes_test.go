package main

import (
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/plans",
	"/subscribe",
}

func Test_RoutesExist(t *testing.T) {
	testRoutes := testapp.Routes()

	chiRoutes, ok := testRoutes.(chi.Router)
	if !ok {
		t.Error("couldn't cast to chi.router")
		return
	}

	for _, r := range routes {
		routestExist(t, r, chiRoutes)
	}

}

func routestExist(t *testing.T, findRoute string, routes chi.Router) {
	found := false

	_ = chi.Walk(routes,
		func(method, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
			if route == findRoute {
				found = true
			}
			return nil
		},
	)

	if !found {
		t.Errorf("did not find %s route in registered routes", findRoute)
	}
}
