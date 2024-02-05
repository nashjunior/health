package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type ResponseSearchableEnterprises struct {
	Total big.Int
	Items []entities.Enterprise
}

type SearchParamEnterprises struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IEnterprisesRepository interface {
	FindByUUID(id uuid.UUID, tx *any) (*entities.Enterprise, error)
	FindByID(id string, tx *any) (*entities.Enterprise, error)
	//Find(params *SearchParamPharmacies) []entities.Enterprise
	//FindAndCount(params *SearchParamPharmacies) IResponseSearchablePharmacies

	Create(entitiy entities.Enterprise, tx *any) error
	//CreateMany(entitiy []entities.Enterprise) error

	Update(entity entities.Enterprise, tx *any) error

	Delete(entity entities.Enterprise, tx *any) error
	DeleteMany(entidades []entities.Enterprise, tx *any) error
}
