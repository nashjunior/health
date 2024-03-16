package controllers

import (
	"encoding/json"
	createdepartment "health/core/clients/application/use-cases/create-department"
	deletedepartment "health/core/clients/application/use-cases/delete-department"
	finddepartmentmanagers "health/core/clients/application/use-cases/find-deparment-managers"
	finddeparrtment "health/core/clients/application/use-cases/find-deparrtment"
	finddepartmentsubordinates "health/core/clients/application/use-cases/find-department-subordinates"
	"health/core/clients/infra/db/gorm/repositories"
	"net/http"

	"github.com/gorilla/mux"
)

type DepartmentsController struct{}

func (controller *DepartmentsController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := finddeparrtment.New(repositories.NewDepartmentssRepositoryGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *DepartmentsController) FindManagersById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := finddepartmentmanagers.New(repositories.NewDepartmentssRepositoryGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *DepartmentsController) FindSubordinatesById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := finddepartmentsubordinates.New(repositories.NewDepartmentssRepositoryGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *DepartmentsController) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := createdepartment.New(
		repositories.NewDepartmentssRepositoryGorm(),
		repositories.NewDepartmentsHiearchyRepositoryGorm(),
	)

	var idManager *string
	if body["id_manager"] != nil {
		id := body["id_manager"].(string)
		idManager = &id
	}

	output, err := usecase.Execute(createdepartment.Input{
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

func (controller *DepartmentsController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := deletedepartment.New(repositories.NewDepartmentssRepositoryGorm(), repositories.NewDepartmentsHiearchyRepositoryGorm())

	err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return

}
