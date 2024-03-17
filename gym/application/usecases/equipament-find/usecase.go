package equipamentfind

import (
	"health/gym/application/dtos"
	"health/gym/domain/repositories"
)

type Usecase struct {
	equipamentsRepository repositories.IEquipamentsRepository
}

func (usecase *Usecase) Execute(id string) (*dtos.EquipamentOutput, error) {

	disease, err := usecase.equipamentsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.EquipamentOutput{
		Uuid:       disease.GetID().String(),
		Name:       disease.GetName(),
		CreatedAt:  disease.CreatedAt,
		UpdateddAt: disease.UpdatedAt,
		DeletedAt:  disease.DeletedAt,
	}, nil
}

func New(equipamentsRepository repositories.IEquipamentsRepository) *Usecase {
	return &Usecase{equipamentsRepository: equipamentsRepository}
}
