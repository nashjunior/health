package createtransaction

import (
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"health/core/infra/db/gorm"
	"time"
)

type Input struct {
	Date              time.Time
	Value             float64
	IdTransactionType string
	Status            *entities.Status
}

type Output struct {
	Id                string    `json:"id"`
	Date              time.Time `json:"date"`
	Value             float64   `json:"value"`
	IdTransactionType string    `json:"id_transaction_type"`
	Status            uint8     `json:"status"`
	CreatedAt         time.Time `json:"created_at"`
}

type Usecase struct {
	typesTransactionsRepository  repositories.ITypeTransactionsRepository
	transactionsRepository       repositories.ITransactionsRepository
	statusTransactionsRepository repositories.IStatusTransactionsRepository
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {
	typeTransaction, err := usecase.typesTransactionsRepository.FindByID(input.IdTransactionType, nil)

	if err != nil {
		return nil, err
	}

	transaction, err := entities.NewTransaction(&input.Date, &input.Value, typeTransaction, nil)

	if err != nil {
		return nil, err
	}

	if input.Status == nil {
		status := entities.Approved
		input.Status = &status
	}

	statusTransaction, err := entities.NewStatusTransaction(input.Status, transaction, nil)

	if err != nil {
		return nil, err
	}

	gorm.Connection.StartTransaction()

	err = usecase.transactionsRepository.Create(transaction, &gorm.Connection)

	if err != nil {
		gorm.Connection.RollbackTransaction()
		return nil, err
	}

	err = usecase.statusTransactionsRepository.Create(statusTransaction, &gorm.Connection)

	if err != nil {
		gorm.Connection.RollbackTransaction()
		return nil, err
	}

	gorm.Connection.CommitTransaction()

	return &Output{
		Id:                transaction.GetID().String(),
		Date:              transaction.GetDate(),
		Value:             transaction.GetValue(),
		Status:            uint8(statusTransaction.GetStatus()),
		IdTransactionType: input.IdTransactionType,
		CreatedAt:         transaction.CreatedAt,
	}, nil
}

func New(
	typesTransactionsRepository repositories.ITypeTransactionsRepository,
	transactionsRepository repositories.ITransactionsRepository,
	statusTransactionsRepository repositories.IStatusTransactionsRepository,
) Usecase {
	return Usecase{
		typesTransactionsRepository:  typesTransactionsRepository,
		transactionsRepository:       transactionsRepository,
		statusTransactionsRepository: statusTransactionsRepository,
	}
}
