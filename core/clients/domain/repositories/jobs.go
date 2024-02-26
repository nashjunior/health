package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type IResponseSearchableJobs struct {
	Total big.Int
	Items []entities.Job
}

type SearchParamJob struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IJobsRepository interface {
	FindByUUID(id uuid.UUID, tx coreRepository.IConnection) (*entities.Job, error)
	FindByID(id string, tx coreRepository.IConnection) (*entities.Job, error)
	FindByName(name string, tx coreRepository.IConnection) (*entities.Job, error)
	//Find(params *SearchParamJob) []entities.Job
	FindAndCount(params *SearchParamJob, tx coreRepository.IConnection) IResponseSearchableJobs

	Create(entitiy *entities.Job, tx coreRepository.IConnection) error
	//CreateMany(entitiy []entities.Job) error

	Update(entity entities.Job, tx coreRepository.IConnection) error

	Delete(entity entities.Job, tx coreRepository.IConnection) error
	DeleteMany(entidades []entities.Job, tx coreRepository.IConnection) error

	FindManagers(id uuid.UUID, tx coreRepository.IConnection) (*[]entities.Job, error)
	FindManagersByJob(job *entities.Job, tx coreRepository.IConnection) (*entities.Job, error)

	FindSubordinates(id uuid.UUID, tx coreRepository.IConnection) (*[]entities.Job, error)
	FindSubordinatesByJob(job *entities.Job, tx coreRepository.IConnection) (*entities.Job, error)
}
