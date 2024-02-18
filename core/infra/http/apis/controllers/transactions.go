package controllers

import (
	"encoding/json"
	createtransaction "health/core/clients/application/use-cases/create-transaction"
	deletetransaction "health/core/clients/application/use-cases/delete-transaction"
	findtransaction "health/core/clients/application/use-cases/find-transaction"
	"health/core/clients/domain/entities"
	"health/core/infra/db/gorm/repositories"

	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type TransactionsController struct{}

func (controller *TransactionsController) FindById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := findtransaction.New(repositories.NewTransactionsGorm(), repositories.NewStatusTransactionsGorm())

	output, err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusCreated)
	response.Write(jsonResp)
	return

}

func (controller *TransactionsController) Create(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	body := map[string]any{}

	err := json.NewDecoder(request.Body).Decode(&body)

	if err != nil {
		http.Error(response, err.Error(), http.StatusBadRequest)
		return
	}

	usecase := createtransaction.New(
		repositories.NewTypesTransactionsGorm(),
		repositories.NewTransactionsGorm(),
		repositories.NewStatusTransactionsGorm(),
	)

	date, err := time.Parse(time.RFC3339, body["date"].(string))

	if err != nil {
		http.Error(response, "Could not parse date", http.StatusUnprocessableEntity)
	}

	var status *uint8

	if body["status"] != nil {
		state := uint8(body["status"].(float64))
		status = &state
	}

	output, err := usecase.Execute(createtransaction.Input{
		Date:              date,
		Value:             body["value"].(float64),
		IdTransactionType: body["id_transaction_type"].(string),
		Status:            (*entities.Status)(status),
	})

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(&output)

	response.WriteHeader(http.StatusCreated)
	response.Write(jsonResp)
	return

}

func (controller *TransactionsController) Delete(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(request)

	usecase := deletetransaction.New(repositories.NewTransactionsGorm(), repositories.NewStatusTransactionsGorm())

	err := usecase.Execute(vars["id"])

	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
	}

	response.WriteHeader(http.StatusNoContent)
	return

}
