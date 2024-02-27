package finddepartmentsubordinates

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

func (usecase *Usecase) Execute(id string) ([]Output, error) {

	job, err := usecase.departmentsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	departmentWithSubordinates, err := usecase.departmentsRepository.FindSubordinatesByDepartment(job, nil)

	if err != nil {
		return nil, err
	}

	subordinates := departmentWithSubordinates.GetSubordinates()

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

func New(departmentsRepository repositories.IDepartmentsRepository) Usecase {
	return Usecase{departmentsRepository: departmentsRepository}
}
