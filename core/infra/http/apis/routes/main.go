package routers

import (
	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()

	return ConfigureRoutes(r, append(
		UsersRoutes,

		AuthRoutes...,
	))
}
