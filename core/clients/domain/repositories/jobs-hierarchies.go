package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableJobsHirearchy struct {
	Total big.Int
	Items []entities.JobHierarchy
}

type SearchParamJobHierarchy struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IJobsHierarchiesRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.JobHierarchy, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.JobHierarchy, error)
	//Find(params *SearchParamJobHierarchy) []entities.JobHierarchy
	FindAndCount(params *SearchParamJobHierarchy, tx coreRepository.IConnection) IResponseSearchableJobsHirearchy

	Create(entitiy *entities.JobHierarchy, tx coreRepository.IConnection) error
	CreateMany(entitiy []entities.JobHierarchy, tx coreRepository.IConnection) error

	Update(entity entities.JobHierarchy, tx coreRepository.IConnection) error

	Delete(entity entities.JobHierarchy, tx coreRepository.IConnection) error
	DeleteByJob(job entities.Job, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.JobHierarchy, tx coreRepository.IConnection) error
}
