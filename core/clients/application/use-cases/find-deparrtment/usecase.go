package finddeparrtment

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
	departmentsRepository repositories.IDepartmentsRepository
}

func (usecase *Usecase) Execute(id string) (*Output, error) {

	job, err := usecase.departmentsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	return &Output{
		Id:        job.GetID().String(),
		Name:      job.GetName(),
		CreatedAt: job.CreatedAt,
		UpdatedAt: job.UpdatedAt,
	}, nil
}

func New(departmentsRepository repositories.IDepartmentsRepository) Usecase {
	return Usecase{departmentsRepository: departmentsRepository}
}
