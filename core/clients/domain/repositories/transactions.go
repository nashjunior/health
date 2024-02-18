package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableTransactions struct {
	Total big.Int
	Items []entities.Transaction
}

type SearchParamTransactions struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type ITransactionsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Transaction, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Transaction, error)
	//Find(params *SearchParamTransactions) []entities.Transaction
	FindAndCount(params *SearchParamTransactions, tx coreRepository.IConnection) IResponseSearchableTransactions

	FindAllByTypeTransaction(id string, tx coreRepository.IConnection) []entities.Transaction

	Create(entitiy *entities.Transaction, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Transaction) error

	Update(entity entities.Transaction, tx coreRepository.IConnection) error

	Delete(entity entities.Transaction, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Transaction, tx coreRepository.IConnection) error
}
