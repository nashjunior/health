package findjobmanagers

import (
	"health/core/clients/domain/repositories"
	"time"
)

type Input struct {
	Id string
}

type Output struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Usecase struct {
	jobsRepository repositories.IJobsRepository
}

func (usecase *Usecase) Execute(id string) ([]Output, error) {

	job, err := usecase.jobsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	jobWithManager, err := usecase.jobsRepository.FindManagersByJob(job, nil)

	if err != nil {
		return nil, err
	}

	managers := jobWithManager.GetManagers()

	var managersFormated []Output

	for _, item := range managers {
		managersFormated = append(managersFormated, Output{
			Id:        item.GetID().String(),
			Name:      item.GetName(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return managersFormated, nil
}

func New(jobsRepository repositories.IJobsRepository) Usecase {
	return Usecase{jobsRepository: jobsRepository}
}
