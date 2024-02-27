package routers

import (
	"health/core/infra/http/apis/controllers"
	"net/http"
)

var departmentsController = controllers.DepartmentsController{}

var DepartmentsRoutes = []Route{
	{
		URI:                   "/departments/{id}",
		Method:                http.MethodGet,
		Callback:              departmentsController.FindById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/departments/{id}/managers",
		Method:                http.MethodGet,
		Callback:              departmentsController.FindManagersById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/departments/{id}/subordinates",
		Method:                http.MethodGet,
		Callback:              departmentsController.FindSubordinatesById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/departments",
		Method:                http.MethodPost,
		Callback:              departmentsController.Create,
		RequireAuthentication: true,
	},
	{
		URI:                   "/departments/{id}",
		Method:                http.MethodDelete,
		Callback:              departmentsController.Delete,
		RequireAuthentication: true,
	},
}
