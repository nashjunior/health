package controllers

import (
	"encoding/json"

	createdisease "health/health/application/usecases/create-disease"
	deletedisease "health/health/application/usecases/delete-disease"
	finddisease "health/health/application/usecases/find-disease"
	inmemory "health/health/infra/db/in-memory"
	"net/http"

	"github.com/gorilla/mux"
)

type DiseasesController struct {
	findDiseaseUsecase   finddisease.Usecase
	creatediseaseUsecase createdisease.Usecase
	deletediseaseUsecase deletedisease.Usecase
}

func (controller *DiseasesController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	output, err := controller.findDiseaseUsecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *DiseasesController) Create(response http.ResponseWriter, request *http.Request) {
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

	output, err := controller.creatediseaseUsecase.Execute(createdisease.Input{
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

func (controller *DiseasesController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(request)

	err := controller.deletediseaseUsecase.Execute(vars["id"])

	if err != nil {

		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return
}

func NewDiseasesController() DiseasesController {
	repo := inmemory.NewDiseasesInMemoryRepository()

	return DiseasesController{
		findDiseaseUsecase:   *finddisease.New(repo),
		creatediseaseUsecase: *createdisease.New(repo),
		deletediseaseUsecase: *deletedisease.New(repo),
	}
}
