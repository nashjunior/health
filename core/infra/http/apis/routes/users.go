package routers

import (
	"health/core/infra/http/apis/controllers"
	"net/http"
)

var usersController = controllers.UsersController{}

var UsersRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Callback:              usersController.Create,
		RequireAuthentication: false,
	},
}
