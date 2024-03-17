package routers

import (
	"health/core/clients/infra/http/apis/controllers"
	"health/core/seedwork/infra/http/api"

	"net/http"
)

var typesTransactionsController = controllers.TypesTransactionsController{}

var TypesTransasctionsRoutes = []api.Route{
	{
		URI:                   "/types-transactions/{id}",
		Method:                http.MethodGet,
		Callback:              typesTransactionsController.FindById,
		RequireAuthentication: true,
	},
	{
		URI:                   "/types-transactions",
		Method:                http.MethodPost,
		Callback:              typesTransactionsController.Create,
		RequireAuthentication: true,
	},
}
