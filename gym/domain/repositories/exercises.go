package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/gym/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableExercise struct {
	Total big.Int
	Items []entities.Exercise
}

type SearchParamExercise struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IExercisesRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Exercise, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Exercise, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Exercise, error)
	//Find(params *SearchParamExercise) []entities.Exercise
	FindAndCount(params *SearchParamExercise, tx coreRepository.IConnection) IResponseSearchableExercise

	Create(entitiy *entities.Exercise, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Exercise) error

	Update(entity entities.Exercise, tx coreRepository.IConnection) error

	Delete(entity entities.Exercise, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Exercise, tx coreRepository.IConnection) error
}
