package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/gym/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableEquipament struct {
	Total big.Int
	Items []entities.Equipament
}

type SearchParamEquipament struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IEquipamentsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Equipament, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Equipament, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Equipament, error)
	//Find(params *SearchParamEquipament) []entities.Equipament
	FindAndCount(params *SearchParamEquipament, tx coreRepository.IConnection) IResponseSearchableEquipament

	Create(entitiy *entities.Equipament, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Equipament) error

	Update(entity entities.Equipament, tx coreRepository.IConnection) error

	Delete(entity entities.Equipament, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Equipament, tx coreRepository.IConnection) error
}
