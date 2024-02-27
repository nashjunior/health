package createdepartment

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
	departmentsRepository            repositories.IDepartmentsRepository
	departmentsHierarchiesRepository repositories.IDepartmentsHierarchiesRepository
}

func (usecase *Usecase) Execute(input Input) (*Output, error) {

	departmentWithName, err := usecase.departmentsRepository.FindByName(input.Name, nil)

	if errors.Is(err, &domainErrors.NotFoundError{}) {
		return nil, err
	}

	if departmentWithName != nil {
		return nil, errors.New("A department already exists using name " + input.Name)
	}

	department, err := entities.NewDepartment(&input.Name, nil, nil, nil)

	if err != nil {
		return nil, err
	}

	departmentNoManager, err := entities.NewDepartmentHierarchy(department, nil, nil)

	if err != nil {
		return nil, err
	}

	departmentHiearchies := []entities.DepartmentHierarchy{*departmentNoManager}

	if input.IdManager != nil {
		manager, err := usecase.departmentsRepository.FindByID(*input.IdManager, nil)

		if err != nil {
			return nil, err
		}

		departmentManager, err := entities.NewDepartmentHierarchy(department, manager, nil)

		if err != nil {
			return nil, err
		}

		departmentHiearchies = append(departmentHiearchies, *departmentManager)
	}

	err = usecase.departmentsRepository.Create(department, nil)
	if err != nil {
		return nil, err
	}

	err = usecase.departmentsHierarchiesRepository.CreateMany(departmentHiearchies, nil)

	if err != nil {
		return nil, err
	}

	return &Output{
		Id:        department.GetID().String(),
		Name:      department.GetName(),
		Manager:   input.IdManager,
		CreatedAt: department.CreatedAt,
	}, nil
}

func New(
	departmentsRepository repositories.IDepartmentsRepository,
	departmentsHierarchiesRepository repositories.IDepartmentsHierarchiesRepository,
) Usecase {
	return Usecase{
		departmentsRepository:            departmentsRepository,
		departmentsHierarchiesRepository: departmentsHierarchiesRepository,
	}
}
