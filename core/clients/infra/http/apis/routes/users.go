package routers

import (
	"health/core/clients/infra/http/apis/controllers"
	"health/core/seedwork/infra/http/api"

	"net/http"
)

var usersController = controllers.UsersController{}

var UsersRoutes = []api.Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Callback:              usersController.Create,
		RequireAuthentication: false,
	},
}
