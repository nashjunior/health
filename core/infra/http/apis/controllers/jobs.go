package controllers

import (
	"encoding/json"
	createjob "health/core/clients/application/use-cases/create-job"
	deletejob "health/core/clients/application/use-cases/delete-job"
	findjob "health/core/clients/application/use-cases/find-job"
	findjobmanagers "health/core/clients/application/use-cases/find-job-managers"
	findjobsubordinates "health/core/clients/application/use-cases/find-job-subordinates"
	"health/core/infra/db/gorm/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

type JobsController struct{}

func (controller *JobsController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := findjob.New(repositories.NewJobsRepositoryGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *JobsController) FindManagersById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := findjobmanagers.New(repositories.NewJobsRepositoryGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *JobsController) FindSubordinatesById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := findjobsubordinates.New(repositories.NewJobsRepositoryGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *JobsController) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := createjob.New(
		repositories.NewJobsRepositoryGorm(),
		repositories.NewJobsHiearchyRepositoryGorm(),
	)

	var idManager *string
	if body["id_manager"] != nil {
		id := body["id_manager"].(string)
		idManager = &id
	}

	output, err := usecase.Execute(createjob.Input{
		Name:      body["name"].(string),
		IdManager: idManager,
	})

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusCreated)
	response.Write(jsonResp)
	return
}

func (controller *JobsController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := deletejob.New(repositories.NewJobsRepositoryGorm(), repositories.NewJobsHiearchyRepositoryGorm())

	err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return

}
