package routes

import (
	"health/core/seedwork/infra/http/api"
	"health/gym/infra/http/apis/controllers"
	"net/http"
)

var exercisesController = controllers.NewExercisesController()

var ExercisesRoutes = []api.Route{
	{
		URI:                   "/exercises/{id}",
		Method:                http.MethodGet,
		Callback:              exercisesController.FindById,
		RequireAuthentication: false,
	},
	{
		URI:                   "/exercises",
		Method:                http.MethodPost,
		Callback:              exercisesController.Create,
		RequireAuthentication: false,
	},
	{
		URI:                   "/exercises/{id}",
		Method:                http.MethodDelete,
		Callback:              exercisesController.Delete,
		RequireAuthentication: false,
	},
}
