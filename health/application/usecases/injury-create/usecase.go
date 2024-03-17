package injurycreate

import (
	"errors"
	"health/health/application/dtos"
	"health/health/domain/entities"
	"health/health/domain/repositories"
)

type Input struct {
	Name *string
}

type Usecase struct {
	injuriesRepository repositories.IInjuriesRepository
}

func (usecase *Usecase) Execute(input Input) (*dtos.InjuryOutput, error) {

	previousInjury, _ := usecase.injuriesRepository.FindByName(*input.Name, nil)

	if previousInjury != nil {
		return nil, errors.New("A injury with this name already exists")
	}

	injury, err := entities.NewInjury(input.Name, nil, nil)

	if err != nil {
		return nil, err
	}

	err = usecase.injuriesRepository.Create(injury, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.InjuryOutput{
		Uuid:      injury.GetID().String(),
		Name:      injury.GetName(),
		CreatedAt: injury.CreatedAt,
	}, nil
}

func New(injuriesRepository repositories.IInjuriesRepository) *Usecase {
	return &Usecase{injuriesRepository: injuriesRepository}
}
