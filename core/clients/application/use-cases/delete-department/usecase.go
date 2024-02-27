package deletedepartment

import (
	"health/core/clients/domain/repositories"
)

type Input struct {
	Id string
}

type Usecase struct {
	departmentsRepository            repositories.IDepartmentsRepository
	departmentsHierarchiesRepository repositories.IDepartmentsHierarchiesRepository
}

func (usecase *Usecase) Execute(id string) error {

	deparment, err := usecase.departmentsRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	usecase.departmentsHierarchiesRepository.DeleteByDepartment(*deparment, nil)
	usecase.departmentsRepository.Delete(*deparment, nil)

	return nil
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
