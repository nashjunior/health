package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableStatusTransactions struct {
	Total big.Int
	Items []entities.StatusTransaction
}

type SearchParamStatusTransactions struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IStatusTransactionsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.StatusTransaction, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.StatusTransaction, error)
	//Find(params *SearchParamStatusTransactions) []entities.StatusTransaction
	FindAndCount(params *SearchParamStatusTransactions, tx coreRepository.IConnection) IResponseSearchableStatusTransactions

	FindAllByTransaction(id string, tx coreRepository.IConnection) []entities.StatusTransaction
	FindACurrentStatusTransaction(id string, tx coreRepository.IConnection) (*entities.StatusTransaction, error)

	Create(entitiy *entities.StatusTransaction, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.StatusTransaction) error

	Update(entity entities.StatusTransaction, tx coreRepository.IConnection) error

	Delete(entity entities.StatusTransaction, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.StatusTransaction, tx coreRepository.IConnection) error
}
