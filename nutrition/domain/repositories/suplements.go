package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/nutrition/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableSuplement struct {
	Total big.Int
	Items []entities.Suplement
}

type SearchParamEquipament struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type ISuplementsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Suplement, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Suplement, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Suplement, error)
	//Find(params *SearchParamEquipament) []entities.Suplement
	FindAndCount(params *SearchParamEquipament, tx coreRepository.IConnection) IResponseSearchableSuplement

	Create(entitiy *entities.Suplement, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Suplement) error

	Update(entity entities.Suplement, tx coreRepository.IConnection) error

	Delete(entity entities.Suplement, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Suplement, tx coreRepository.IConnection) error
}
