package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/health/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableDiseases struct {
	Total big.Int
	Items []entities.Disease
}

type SearchParamDisease struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IDiseasesRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Disease, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Disease, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Disease, error)
	//Find(params *SearchParamDisease) []entities.Disease
	FindAndCount(params *SearchParamDisease, tx coreRepository.IConnection) IResponseSearchableDiseases

	Create(entitiy *entities.Disease, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Disease) error

	Update(entity entities.Disease, tx coreRepository.IConnection) error

	Delete(entity entities.Disease, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Disease, tx coreRepository.IConnection) error
}
