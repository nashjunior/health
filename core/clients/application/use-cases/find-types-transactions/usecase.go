package findtypestransactions

import (
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"math/big"
	"time"
)

type Input struct {
}

type Usecase struct {
	typesTransactionsRepository repositories.ITypeTransactionsRepository
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {
	medicaments := usecase.typesTransactionsRepository.FindAndCount(nil, nil)

	var items []ITypeTransactionOutput

	for _, medicament := range medicaments.Items {
		dto := ITypeTransactionOutput{
			Id:        medicament.GetID().String(),
			Name:      medicament.GetName(),
			createdAt: medicament.CreatedAt,
			updatedAt: medicament.UpdatedAt,
		}
		items = append(items, dto)
	}

	return &Output{
		Items:     items,
		Total:     medicaments.Total,
		TotalPage: *big.NewInt(1),
	}, nil
}

type ITypeTransactionOutput struct {
	Id            string                 `json:"id"`
	Name          string                 `json:"name"`
	AccountType   entities.AccountType   `json:"account_type"`
	OperationType entities.OperationType `json:"operation_type"`
	createdAt     time.Time
	updatedAt     *time.Time
}

type Output struct {
	Items     []ITypeTransactionOutput `json:"items"`
	Total     big.Int                  `json:"total"`
	TotalPage big.Int                  `json:"total_page"`
}

func New(typesTransactionsRepository repositories.ITypeTransactionsRepository) Usecase {
	return Usecase{typesTransactionsRepository: typesTransactionsRepository}
}
