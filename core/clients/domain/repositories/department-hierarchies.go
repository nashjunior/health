package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableDepartmentsHirearchy struct {
	Total big.Int
	Items []entities.DepartmentHierarchy
}

type SearchParamDepartmentHierarchy struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IDepartmentsHierarchiesRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.DepartmentHierarchy, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.DepartmentHierarchy, error)
	//Find(params *SearchParamDepartmentHierarchy) []entities.DepartmentHierarchy
	FindAndCount(params *SearchParamDepartmentHierarchy, tx coreRepository.IConnection) IResponseSearchableDepartmentsHirearchy

	Create(entitiy *entities.DepartmentHierarchy, tx coreRepository.IConnection) error
	CreateMany(entitiy []entities.DepartmentHierarchy, tx coreRepository.IConnection) error

	Update(entity entities.DepartmentHierarchy, tx coreRepository.IConnection) error

	Delete(entity entities.DepartmentHierarchy, tx coreRepository.IConnection) error
	DeleteByDepartment(job entities.Department, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.DepartmentHierarchy, tx coreRepository.IConnection) error
}
