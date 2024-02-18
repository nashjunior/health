package createtypetransaction

import (
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"time"
)

type Input struct {
	Name          string
	AccountType   entities.AccountType
	OperationType entities.OperationType
}

type Output struct {
	Id            string                 `json:"id"`
	Name          string                 `json:"name"`
	AccountType   entities.AccountType   `json:"account_type"`
	OperationType entities.OperationType `json:"operation_type"`
	CreatedAt     time.Time              `json:"created_at"`
}

type Usecase struct {
	typesTransactionsRepository repositories.ITypeTransactionsRepository
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {
	transactionType, err := entities.NewTypeTransaction(&input.Name, &input.AccountType, &input.OperationType, nil)

	if err != nil {
		return nil, err
	}

	err = usecase.typesTransactionsRepository.Create(transactionType, nil)

	if err != nil {
		return nil, err
	}

	return &Output{
		Id:            transactionType.GetID().String(),
		Name:          transactionType.GetName(),
		AccountType:   transactionType.GetAccountType(),
		OperationType: transactionType.GetOperationType(),
		CreatedAt:     transactionType.CreatedAt,
	}, nil
}

func New(
	typesTransactionsRepository repositories.ITypeTransactionsRepository,
) Usecase {
	return Usecase{
		typesTransactionsRepository: typesTransactionsRepository,
	}
}
