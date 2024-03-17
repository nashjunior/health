package routes

import (
	"health/core/seedwork/infra/http/api"
	"health/health/infra/http/apis/controllers"
	"net/http"
)

var injuriesController = controllers.NewInjuriesController()

var InjuriesRoutes = []api.Route{
	{
		URI:                   "/injuries/{id}",
		Method:                http.MethodGet,
		Callback:              injuriesController.FindById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/injuries",
		Method:                http.MethodPost,
		Callback:              injuriesController.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/injuries/{id}",
		Method:                http.MethodDelete,
		Callback:              injuriesController.Delete,
		RequireAuthentication: false,
	},
}
