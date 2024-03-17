package controllers

import (
	"encoding/json"

	equipamentcreate "health/gym/application/usecases/equipament-create"
	equipamentdelete "health/gym/application/usecases/equipament-delete"
	equipamentfind "health/gym/application/usecases/equipament-find"
	inmemory "health/gym/infra/db/in-memory"
	"net/http"

	"github.com/gorilla/mux"
)

type EquipamentsController struct {
	findEquipamentUsecase   equipamentfind.Usecase
	createEquipamentUsecase equipamentcreate.Usecase
	deleteEquipamentUsecase equipamentdelete.Usecase
}

func (controller *EquipamentsController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	output, err := controller.findEquipamentUsecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusOK)
	response.Write(jsonResp)
	return

}

func (controller *EquipamentsController) Create(response http.ResponseWriter, request *http.Request) {
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

	output, err := controller.createEquipamentUsecase.Execute(equipamentcreate.Input{
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

func (controller *EquipamentsController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/json")

	vars := mux.Vars(request)

	err := controller.deleteEquipamentUsecase.Execute(vars["id"])

	if err != nil {

		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return
}

func NewEquipamentsController() EquipamentsController {
	repo := inmemory.NewEquipamentsInMemoryRepository()

	return EquipamentsController{
		findEquipamentUsecase:   *equipamentfind.New(repo),
		createEquipamentUsecase: *equipamentcreate.New(repo),
		deleteEquipamentUsecase: *equipamentdelete.New(repo),
	}
}
