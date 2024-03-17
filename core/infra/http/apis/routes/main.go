package routes

import (
	routersCore "health/core/clients/infra/http/apis/routes"
	"slices"

	routersHealth "health/health/infra/http/apis/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()

	codeRoute := routersCore.AuthRoutes[0]

	r.HandleFunc(codeRoute.URI, codeRoute.Callback)

	return ConfigureRoutes(
		r, slices.Concat(
			routersCore.AuthRoutes[1:],
			routersCore.JobsRoutes,
			routersCore.DepartmentsRoutes,
			routersCore.TypesTransasctionsRoutes,
			routersCore.TransasctionsRoutes,
			routersHealth.DiseaseRoutes,
		),
	)

}
