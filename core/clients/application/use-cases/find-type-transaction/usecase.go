package findtypetransaction

import (
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"time"
)

type Usecase struct {
	typesTransactionsRepository repositories.ITypeTransactionsRepository
}

func (usecase *Usecase) Execute(id string) (*Output, error) {
	typeTransaction, err := usecase.typesTransactionsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	return &Output{
		Id:            typeTransaction.GetID().String(),
		Name:          typeTransaction.GetName(),
		AccountType:   typeTransaction.GetAccountType(),
		OperationType: typeTransaction.GetOperationType(),
		createdAt:     typeTransaction.CreatedAt,
		updatedAt:     typeTransaction.UpdatedAt,
	}, err
}

type Output struct {
	Id            string                 `json:"id"`
	Name          string                 `json:"name"`
	AccountType   entities.AccountType   `json:"account_type"`
	OperationType entities.OperationType `json:"operation_type"`
	createdAt     time.Time
	updatedAt     *time.Time
}

func New(typesTransactionsRepository repositories.ITypeTransactionsRepository) Usecase {
	return Usecase{typesTransactionsRepository: typesTransactionsRepository}
}
