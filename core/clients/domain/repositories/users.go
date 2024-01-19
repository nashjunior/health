package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableUsers struct {
	Total big.Int
	Items []entities.User
}

type SearchParamUsers struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IUsersRepository interface {
	FindByUUID(id uuid.UUID, tx *any) (*entities.User, error)
	FindByID(id string, tx *any) (*entities.User, error)
	//Find(params *SearchParamUsers) []entities.User
	//FindAndCount(params *SearchParamUsers) IResponseSearchableUsers

	Create(entitiy entities.User, tx *any) error
	//CreateMany(entitiy []entities.User) error

	Update(entity entities.User, tx *any) error

	Delete(entity entities.User, tx *any) error
	DeleteMany(entidades []entities.User, tx *any) error
}
