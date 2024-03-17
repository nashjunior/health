package routers

import (
	"health/core/clients/infra/http/apis/controllers"
	"health/core/seedwork/infra/http/api"

	"net/http"
)

var authController = controllers.AuthController{}

var AuthRoutes = []api.Route{
	{
		URI:                   "/code",
		Method:                http.MethodPost,
		Callback:              authController.GetCode,
		RequireAuthentication: false,
	},
	{
		URI:                   "/refresh-token",
		Method:                http.MethodPost,
		Callback:              authController.RefreshToken,
		RequireAuthentication: false,
	},
	{
		URI:                   "/token-exchange",
		Method:                http.MethodPost,
		Callback:              authController.ExchangeAccessToken,
		RequireAuthentication: false,
	},
	{
		URI:                   "/sign-in",
		Method:                http.MethodPost,
		Callback:              authController.SignIn,
		RequireAuthentication: false,
	},
}
