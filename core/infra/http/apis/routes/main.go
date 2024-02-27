package routers

import (
	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()

	codeRoute := AuthRoutes[0]

	r.HandleFunc(codeRoute.URI, codeRoute.Callback)

	return ConfigureRoutes(r, append(
		AuthRoutes[1:],
		append(
			UsersRoutes,
			append(
				JobsRoutes,
				append(
					DepartmentsRoutes,
					append(
						TypesTransasctionsRoutes,
						TransasctionsRoutes...,
					)...,
				)...,
			)...,
		)...,
	))
}
