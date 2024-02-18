package controllers

import (
	"encoding/json"
	"fmt"
	createtypetransaction "health/core/clients/application/use-cases/create-type-transaction"
	findtypetransaction "health/core/clients/application/use-cases/find-type-transaction"
	"health/core/clients/domain/entities"
	"health/core/infra/db/gorm/repositories"

	"net/http"

	"github.com/gorilla/mux"
)

type TypesTransactionsController struct{}

func (controller *TypesTransactionsController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := findtypetransaction.New(repositories.NewTypesTransactionsGorm())

	fmt.Println(vars["id"])

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusCreated)
	response.Write(jsonResp)
	return

}

func (controller *TypesTransactionsController) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := createtypetransaction.New(repositories.NewTypesTransactionsGorm())

	output, err := usecase.Execute(createtypetransaction.Input{
		Name:          body["name"].(string),
		AccountType:   entities.AccountType(body["account_type"].(float64)),
		OperationType: entities.OperationType(body["operation_type"].(float64)),
	})

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusCreated)
	response.Write(jsonResp)
	return

}
