package deletetypetransaction

import (
	"health/core/clients/domain/repositories"
)

type Usecase struct {
	typesTransactionsRepository repositories.ITypeTransactionsRepository
}

func (usecase *Usecase) Execute(id string) error {
	typeTransaction, err := usecase.typesTransactionsRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	return usecase.typesTransactionsRepository.Delete(*typeTransaction, nil)
}

func New(personsRepository repositories.ITypeTransactionsRepository) Usecase {
	return Usecase{typesTransactionsRepository: personsRepository}
}
