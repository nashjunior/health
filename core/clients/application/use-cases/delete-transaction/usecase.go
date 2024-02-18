package deletetransaction

import (
	"health/core/clients/domain/repositories"
	"health/core/infra/db/gorm"
)

type Usecase struct {
	transactionsRepository       repositories.ITransactionsRepository
	statusTransactionsRepository repositories.IStatusTransactionsRepository
}

func (usecase *Usecase) Execute(id string) error {
	transaction, err := usecase.transactionsRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	gorm.Connection.StartTransaction()

	err = usecase.statusTransactionsRepository.DeleteByTransaction(*transaction, nil)

	if err != nil {
		gorm.Connection.RollbackTransaction()
		return err
	}

	err = usecase.transactionsRepository.Delete(*transaction, nil)
	if err != nil {
		gorm.Connection.RollbackTransaction()
		return err
	}

	gorm.Connection.CommitTransaction()

	return nil
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
