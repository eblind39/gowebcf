package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"testing"
)

var routes = []string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/members/plans",
	"/members/subscribe",
}

func Test_Routes_Exist(t *testing.T) {
	testRoutes := testApp.routes()

	chiRoutes := testRoutes.(chi.Router)

	for _, route := range routes {
		routeExists(t, chiRoutes, route)
	}
}

func routeExists(t *testing.T, routes chi.Router, route string) {
	found := false

	_ = chi.Walk(routes, func(method string, foundRoute string, handler http.Handler, middlewares ...func(handler2 http.Handler) http.Handler) error {
		if route == foundRoute {
			found = true
		}

		return nil
	})

	if !found {
		t.Errorf("did not %s in registered routes", route)
	}
}
