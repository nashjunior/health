package repositories

import (
	"health/core/clients/domain/repositories"
	inmemory "health/core/infra/db/in-memory"
)

type Transaction interface {
	StartTransaction() error
	CommiTransaction() error
	RollbackTransaction() error
}

func NewUsersRepository() repositories.IUsersRepository {
	return inmemory.NewUsersInMemoryRepository()
}

func NewPersonsRepository() repositories.IPersonsRepository {
	return inmemory.NewPersonsInMemoryRepository()
}

func NewEnterprisesRepository() repositories.IEnterprisesRepository {
	return inmemory.NewEnterprisesInMemoryRepository()
}

func NewShareholdersRepository() repositories.IShareholdersRepository {
	return inmemory.NewShareholdersInMemoryRepository()
}

func NewValidationsCodesRepository() repositories.IValidationsCodesRepository {
	return inmemory.NewValidationsCodesInMemoryRepository()
}

func NewMedicamentsRepository() repositories.IMedicamentsRepository {
	return inmemory.NewMedicamentsInMemoryRepository()
}

func NewTransaction() Transaction {
	return &inmemory.Transaction{}
}
