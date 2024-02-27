package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableDepartments struct {
	Total big.Int
	Items []entities.Department
}

type SearchParamDepartment struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IDepartmentsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Department, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Department, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Department, error)
	//Find(params *SearchParamDepartment) []entities.Department
	FindAndCount(params *SearchParamDepartment, tx coreRepository.IConnection) IResponseSearchableDepartments

	Create(entitiy *entities.Department, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Department) error

	Update(entity entities.Department, tx coreRepository.IConnection) error

	Delete(entity entities.Department, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Department, tx coreRepository.IConnection) error

	FindManagers(id uuid.UUID, tx coreRepository.IConnection) (*[]entities.Department, error)
	FindManagersByDepartment(department *entities.Department, tx coreRepository.IConnection) (*entities.Department, error)

	FindSubordinates(id uuid.UUID, tx coreRepository.IConnection) (*[]entities.Department, error)
	FindSubordinatesByDepartment(department *entities.Department, tx coreRepository.IConnection) (*entities.Department, error)
}
