package routes

import (
	"health/core/infra/http/apis/middlewares"
	"health/core/seedwork/infra/http/api"

	"github.com/gorilla/mux"
)

func ConfigureRoutes(r *mux.Router, routes []api.Route) *mux.Router {
	for _, route := range routes {
		request := route.Callback
		if route.RequireAuthentication {
			request = middlewares.EnsureAuthenticated(route.Callback)
		}

		r.HandleFunc(route.URI, request).Methods(route.Method)
	}

	return r
}
