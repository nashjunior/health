package controllers

import (
	"encoding/json"

	exercisecreate "health/gym/application/usecases/exercise-create"
	exercisedelete "health/gym/application/usecases/exercise-delete"
	exercisefind "health/gym/application/usecases/exercise-find"
	inmemory "health/gym/infra/db/in-memory"
	"net/http"

	"github.com/gorilla/mux"
)

type ExercisesController struct {
	findExerciseUsecase   exercisefind.Usecase
	createExerciseUsecase exercisecreate.Usecase
	deleteExerciseUsecase exercisedelete.Usecase
}

func (controller *ExercisesController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	output, err := controller.findExerciseUsecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *ExercisesController) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	var name *string

	if body["name"] != nil {
		convertion := body["name"].(string)
		name = &convertion
	}

	output, err := controller.createExerciseUsecase.Execute(exercisecreate.Input{
		Name: name,
	})

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusCreated)
	response.Write(jsonResp)
	return
}

func (controller *ExercisesController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(request)

	err := controller.deleteExerciseUsecase.Execute(vars["id"])

	if err != nil {

		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return
}

func NewExercisesController() ExercisesController {
	repo := inmemory.NewExercisesInMemoryRepository()

	return ExercisesController{
		findExerciseUsecase:   *exercisefind.New(repo),
		createExerciseUsecase: *exercisecreate.New(repo),
		deleteExerciseUsecase: *exercisedelete.New(repo),
	}
}
