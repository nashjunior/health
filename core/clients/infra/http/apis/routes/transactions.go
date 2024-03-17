package routers

import (
	"health/core/clients/infra/http/apis/controllers"

	"health/core/seedwork/infra/http/api"
	"net/http"
)

var transactionsController = controllers.TransactionsController{}

var TransasctionsRoutes = []api.Route{
	{
		URI:                   "/transactions/{id}",
		Method:                http.MethodGet,
		Callback:              transactionsController.FindById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/transactions",
		Method:                http.MethodPost,
		Callback:              transactionsController.Create,
		RequireAuthentication: true,
	},
	{
		URI:                   "/transactions/{id}",
		Method:                http.MethodDelete,
		Callback:              transactionsController.Delete,
		RequireAuthentication: true,
	},
}
