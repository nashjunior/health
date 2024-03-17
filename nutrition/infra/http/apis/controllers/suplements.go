package controllers

import (
	"encoding/json"

	suplementcreate "health/nutrition/application/usecases/suplement-create"
	suplementdelete "health/nutrition/application/usecases/suplement-delete"
	suplementfind "health/nutrition/application/usecases/suplement-find"
	inmemory "health/nutrition/infra/db/in-memory"
	"net/http"

	"github.com/gorilla/mux"
)

type SuplementsController struct {
	findSuplementUsecase   suplementfind.Usecase
	createSuplementUsecase suplementcreate.Usecase
	deleteSuplementUsecase suplementdelete.Usecase
}

func (controller *SuplementsController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	output, err := controller.findSuplementUsecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *SuplementsController) Create(response http.ResponseWriter, request *http.Request) {
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

	output, err := controller.createSuplementUsecase.Execute(suplementcreate.Input{
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

func (controller *SuplementsController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(request)

	err := controller.deleteSuplementUsecase.Execute(vars["id"])

	if err != nil {

		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return
}

func NewSuplementsController() SuplementsController {
	repo := inmemory.NewSuplementsInMemoryRepository()

	return SuplementsController{
		findSuplementUsecase:   *suplementfind.New(repo),
		createSuplementUsecase: *suplementcreate.New(repo),
		deleteSuplementUsecase: *suplementdelete.New(repo),
	}
}
