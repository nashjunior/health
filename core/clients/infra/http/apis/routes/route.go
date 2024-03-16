package routers

import (
	"health/core/infra/http/apis/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                   string
	Method                string
	Callback              func(http.ResponseWriter, *http.Request)
	RequireAuthentication bool
}

func ConfigureRoutes(r *mux.Router, routes []Route) *mux.Router {
	for _, route := range routes {
		request := route.Callback
		if route.RequireAuthentication {
			request = middlewares.EnsureAuthenticated(route.Callback)
		}

		r.HandleFunc(route.URI, request).Methods(route.Method)
	}

	return r
}
