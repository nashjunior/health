package injuryfind

import (
	"health/health/application/dtos"
	"health/health/domain/repositories"
)

type Usecase struct {
	injuriesRepository repositories.IInjuriesRepository
}

func (usecase *Usecase) Execute(id string) (*dtos.InjuryOutput, error) {

	injury, err := usecase.injuriesRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.InjuryOutput{
		Uuid:       injury.GetID().String(),
		Name:       injury.GetName(),
		CreatedAt:  injury.CreatedAt,
		UpdateddAt: injury.UpdatedAt,
		DeletedAt:  injury.DeletedAt,
	}, nil
}

func New(injuriesRepository repositories.IInjuriesRepository) *Usecase {
	return &Usecase{injuriesRepository: injuriesRepository}
}
