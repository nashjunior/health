package controllers

import (
	"encoding/json"

	injurycreate "health/health/application/usecases/injury-create"
	injurydelete "health/health/application/usecases/injury-delete"
	injuryfind "health/health/application/usecases/injury-find"

	inmemory "health/health/infra/db/in-memory"
	"net/http"

	"github.com/gorilla/mux"
)

type InjuriesController struct {
	findInjuryUsecase   injuryfind.Usecase
	createInjuryUsecase injurycreate.Usecase
	deleteInjuryUsecase injurydelete.Usecase
}

func (controller *InjuriesController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	output, err := controller.findInjuryUsecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *InjuriesController) Create(response http.ResponseWriter, request *http.Request) {
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

	output, err := controller.createInjuryUsecase.Execute(injurycreate.Input{
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

func (controller *InjuriesController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(request)

	err := controller.deleteInjuryUsecase.Execute(vars["id"])

	if err != nil {

		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return
}

func NewInjuriesController() InjuriesController {
	repo := inmemory.NewInjuriesInMemoryRepository()

	return InjuriesController{
		findInjuryUsecase:   *injuryfind.New(repo),
		createInjuryUsecase: *injurycreate.New(repo),
		deleteInjuryUsecase: *injurydelete.New(repo),
	}
}
