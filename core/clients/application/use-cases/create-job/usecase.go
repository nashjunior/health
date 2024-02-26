package createjob

import (
	"errors"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"time"

	domainErrors "health/core/application/errors"
)

type Input struct {
	Name      string
	IdManager *string
}

type Output struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Manager   *string   `json:"manager"`
	CreatedAt time.Time `json:"created_at"`
}

type Usecase struct {
	jobsRepository            repositories.IJobsRepository
	jobsHierarchiesRepository repositories.IJobsHierarchiesRepository
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {

	jobWithName, err := usecase.jobsRepository.FindByName(input.Name, nil)

	if errors.Is(err, &domainErrors.NotFoundError{}) {
		return nil, err
	}

	if jobWithName != nil {
		return nil, errors.New("A job already exists using name " + input.Name)
	}

	job, err := entities.NewJob(&input.Name, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	jobNoManager, err := entities.NewJobHierarchy(job, nil, nil)

	if err != nil {
		return nil, err
	}

	jobHiearchies := []entities.JobHierarchy{*jobNoManager}

	if input.IdManager != nil {
		manager, err := usecase.jobsRepository.FindByID(*input.IdManager, nil)

		if err != nil {
			return nil, err
		}

		jobManager, err := entities.NewJobHierarchy(job, manager, nil)

		if err != nil {
			return nil, err
		}

		jobHiearchies = append(jobHiearchies, *jobManager)
	}

	err = usecase.jobsRepository.Create(job, nil)
	if err != nil {
		return nil, err
	}

	err = usecase.jobsHierarchiesRepository.CreateMany(jobHiearchies, nil)

	if err != nil {
		return nil, err
	}

	return &Output{
		Id:        job.GetID().String(),
		Name:      job.GetName(),
		Manager:   input.IdManager,
		CreatedAt: job.CreatedAt,
	}, nil
}

func New(
	jobsRepository repositories.IJobsRepository,
	jobsHierarchiesRepository repositories.IJobsHierarchiesRepository,
) Usecase {
	return Usecase{
		jobsRepository:            jobsRepository,
		jobsHierarchiesRepository: jobsHierarchiesRepository,
	}
}
