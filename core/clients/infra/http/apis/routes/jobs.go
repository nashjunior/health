package routers

import (
	"health/core/clients/infra/http/apis/controllers"
	"net/http"
)

var jobsController = controllers.JobsController{}

var JobsRoutes = []Route{
	{
		URI:                   "/jobs/{id}",
		Method:                http.MethodGet,
		Callback:              jobsController.FindById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/jobs/{id}/managers",
		Method:                http.MethodGet,
		Callback:              jobsController.FindManagersById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/jobs/{id}/subordinates",
		Method:                http.MethodGet,
		Callback:              jobsController.FindSubordinatesById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/jobs",
		Method:                http.MethodPost,
		Callback:              jobsController.Create,
		RequireAuthentication: true,
	},
	{
		URI:                   "/jobs/{id}",
		Method:                http.MethodDelete,
		Callback:              jobsController.Delete,
		RequireAuthentication: true,
	},
}
