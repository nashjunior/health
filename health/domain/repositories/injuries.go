package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/health/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableInjury struct {
	Total big.Int
	Items []entities.Injury
}

type SearchParamInjury struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IInjuriesRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Injury, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Injury, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Injury, error)
	//Find(params *SearchParamInjury) []entities.Injury
	FindAndCount(params *SearchParamInjury, tx coreRepository.IConnection) IResponseSearchableInjury

	Create(entitiy *entities.Injury, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Injury) error

	Update(entity entities.Injury, tx coreRepository.IConnection) error

	Delete(entity entities.Injury, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Injury, tx coreRepository.IConnection) error
}
