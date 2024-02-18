package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableTypeTransactions struct {
	Total big.Int
	Items []entities.TypeTransaction
}

type SearchParamTypeTransactions struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type ITypeTransactionsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.TypeTransaction, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.TypeTransaction, error)
	//Find(params *SearchParamTypeTransactions) []entities.TypeTransaction
	FindAndCount(params *SearchParamTypeTransactions, tx coreRepository.IConnection) IResponseSearchableTypeTransactions

	Create(entitiy *entities.TypeTransaction, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.TypeTransaction) error

	Update(entity entities.TypeTransaction, tx coreRepository.IConnection) error

	Delete(entity entities.TypeTransaction, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.TypeTransaction, tx coreRepository.IConnection) error
}
