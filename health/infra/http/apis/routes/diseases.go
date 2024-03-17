package routes

import (
	"health/core/seedwork/infra/http/api"
	"health/health/infra/http/apis/controllers"
	"net/http"
)

var diseasesController = controllers.NewDiseasesController()

var DiseaseRoutes = []api.Route{
	{
		URI:                   "/diseases/{id}",
		Method:                http.MethodGet,
		Callback:              diseasesController.FindById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/diseases",
		Method:                http.MethodPost,
		Callback:              diseasesController.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/diseases/{id}",
		Method:                http.MethodDelete,
		Callback:              diseasesController.Delete,
		RequireAuthentication: false,
	},
}
