package routes

import (
	"health/core/seedwork/infra/http/api"
	"health/gym/infra/http/apis/controllers"
	"net/http"
)

var equipamentsController = controllers.NewEquipamentsController()

var EquipamentsRoutes = []api.Route{
	{
		URI:                   "/equipaments/{id}",
		Method:                http.MethodGet,
		Callback:              equipamentsController.FindById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/equipaments",
		Method:                http.MethodPost,
		Callback:              equipamentsController.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/equipaments/{id}",
		Method:                http.MethodDelete,
		Callback:              equipamentsController.Delete,
		RequireAuthentication: false,
	},
}
