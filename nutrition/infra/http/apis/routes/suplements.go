package routes

import (
	"health/core/seedwork/infra/http/api"
	"health/nutrition/infra/http/apis/controllers"
	"net/http"
)

var suplementsController = controllers.NewSuplementsController()

var SuplmentsRoutes = []api.Route{
	{
		URI:                   "/suplements/{id}",
		Method:                http.MethodGet,
		Callback:              suplementsController.FindById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/suplements",
		Method:                http.MethodPost,
		Callback:              suplementsController.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/suplements/{id}",
		Method:                http.MethodDelete,
		Callback:              suplementsController.Delete,
		RequireAuthentication: false,
	},
}
