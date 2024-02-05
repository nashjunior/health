package routers

import (
	"health/core/infra/http/apis/controllers"
	"net/http"
)

var authController = controllers.AuthController{}

var AuthRoutes = []Route{
	{
		URI:                   "/refresh-token",
		Method:                http.MethodPost,
		Callback:              authController.RefreshToken,
		RequireAuthentication: false,
	},
	{
		URI:                   "/sign-in",
		Method:                http.MethodPost,
		Callback:              authController.SignIn,
		RequireAuthentication: false,
	},
}
