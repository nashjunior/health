package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type ResponseSearchablePersons struct {
	Total big.Int
	Items []entities.Person
}

type SearchParamPersons struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IPersonsRepository interface {
	FindByUUID(id uuid.UUID, tx *any) (*entities.Person, error)
	FindByID(id string, tx *any) (*entities.Person, error)
	//Find(params *SearchParamPharmacies) []entities.Person
	//FindAndCount(params *SearchParamPharmacies) IResponseSearchablePharmacies

	FindByUser(user entities.User, tx *any) (*entities.Person, error)

	Create(entitiy entities.Person, tx *any) error
	//CreateMany(entitiy []entities.Person, tx *any) error

	Update(entity entities.Person, tx *any) error

	Delete(entity entities.Person, tx *any) error
	DeleteMany(entidades []entities.Person, tx *any) error
}
