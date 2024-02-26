package findjobsubordinates

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

	jobWithManager, err := usecase.jobsRepository.FindSubordinatesByJob(job, nil)

	if err != nil {
		return nil, err
	}

	subordinates := jobWithManager.GetSubordinates()

	var subordinatesFormated []Output

	for _, item := range subordinates {
		subordinatesFormated = append(subordinatesFormated, Output{
			Id:        item.GetID().String(),
			Name:      item.GetName(),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return subordinatesFormated, nil
}

func New(jobsRepository repositories.IJobsRepository) Usecase {
	return Usecase{jobsRepository: jobsRepository}
}
