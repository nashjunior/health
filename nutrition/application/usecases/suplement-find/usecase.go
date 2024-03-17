package suplementfind

import (
	"health/nutrition/application/dtos"
	"health/nutrition/domain/repositories"
)

type Usecase struct {
	suplementsRepository repositories.ISuplementsRepository
}

func (usecase *Usecase) Execute(id string) (*dtos.SuplementOutput, error) {

	disease, err := usecase.suplementsRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.SuplementOutput{
		Uuid:       disease.GetID().String(),
		Name:       disease.GetName(),
		CreatedAt:  disease.CreatedAt,
		UpdateddAt: disease.UpdatedAt,
		DeletedAt:  disease.DeletedAt,
	}, nil
}

func New(suplementsRepository repositories.ISuplementsRepository) *Usecase {
	return &Usecase{suplementsRepository: suplementsRepository}
}
