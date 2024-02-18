package findtransaction

import (
	"health/core/clients/domain/repositories"
	"time"
)

type Usecase struct {
	transactionsRepository       repositories.ITransactionsRepository
	statusTransactionsRepository repositories.IStatusTransactionsRepository
}

func (usecase *Usecase) Execute(id string) (*Output, error) {

	status, err := usecase.statusTransactionsRepository.FindACurrentStatusTransaction(id, nil)

	if err != nil {
		return nil, err
	}

	transaction, err := usecase.transactionsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	typeTransaction := transaction.GetTypeTransaction()

	return &Output{
		Id:                transaction.GetID().String(),
		Date:              transaction.GetDate(),
		Value:             transaction.GetValue(),
		IdTransactionType: typeTransaction.GetID().String(),
		Status:            uint8(status.GetStatus()),
		CreatedAt:         transaction.CreatedAt,
		UpdatedAt:         transaction.UpdatedAt,
	}, err
}

type Output struct {
	Id                string    `json:"id"`
	Date              time.Time `json:"date"`
	Value             float64   `json:"value"`
	IdTransactionType string    `json:"id_transaction_type"`
	Status            uint8     `json:"status"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func New(
	transactionsRepository repositories.ITransactionsRepository,
	statusTransactionsRepository repositories.IStatusTransactionsRepository,
) Usecase {
	return Usecase{
		transactionsRepository:       transactionsRepository,
		statusTransactionsRepository: statusTransactionsRepository,
	}
}
