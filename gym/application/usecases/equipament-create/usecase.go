package equipamentcreate

import (
	"errors"
	"health/gym/application/dtos"
	"health/gym/domain/entities"
	"health/gym/domain/repositories"
)

type Input struct {
	Name *string
}

type Usecase struct {
	equipamentsRepository repositories.IEquipamentsRepository
}

func (usecase *Usecase) Execute(input Input) (*dtos.EquipamentOutput, error) {

	previousEquipament, _ := usecase.equipamentsRepository.FindByName(*input.Name, nil)

	if previousEquipament != nil {
		return nil, errors.New("A equipament with this name already exists")
	}

	equipament, err := entities.NewEquipament(input.Name, nil, nil)

	if err != nil {
		return nil, err
	}

	err = usecase.equipamentsRepository.Create(equipament, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.EquipamentOutput{
		Uuid:      equipament.GetID().String(),
		Name:      equipament.GetName(),
		CreatedAt: equipament.CreatedAt,
	}, nil
}

func New(equipamentsRepository repositories.IEquipamentsRepository) *Usecase {
	return &Usecase{equipamentsRepository: equipamentsRepository}
}
