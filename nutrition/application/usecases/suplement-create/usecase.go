package suplementcreate

import (
	"errors"
	"health/nutrition/application/dtos"
	"health/nutrition/domain/entities"
	"health/nutrition/domain/repositories"
)

type Input struct {
	Name *string
}

type Usecase struct {
	suplementsRepository repositories.ISuplementsRepository
}

func (usecase *Usecase) Execute(input Input) (*dtos.SuplementOutput, error) {

	previousEquipament, _ := usecase.suplementsRepository.FindByName(*input.Name, nil)

	if previousEquipament != nil {
		return nil, errors.New("A suplement with this name already exists")
	}

	equipament, err := entities.NewSuplement(input.Name, nil, nil)

	if err != nil {
		return nil, err
	}

	err = usecase.suplementsRepository.Create(equipament, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.SuplementOutput{
		Uuid:      equipament.GetID().String(),
		Name:      equipament.GetName(),
		CreatedAt: equipament.CreatedAt,
	}, nil
}

func New(suplementsRepository repositories.ISuplementsRepository) *Usecase {
	return &Usecase{suplementsRepository: suplementsRepository}
}
