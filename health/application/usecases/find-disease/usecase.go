package finddisease

import (
	"health/health/application/dtos"
	"health/health/domain/repositories"
)

type Usecase struct {
	diseasesRepository repositories.IDiseasesRepository
}

func (usecase *Usecase) Execute(id string) (*dtos.DiseaseOutput, error) {

	disease, err := usecase.diseasesRepository.FindByID(id, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.DiseaseOutput{
		Uuid:       disease.GetID().String(),
		Name:       disease.GetName(),
		CreatedAt:  disease.CreatedAt,
		UpdateddAt: disease.UpdatedAt,
		DeletedAt:  disease.DeletedAt,
	}, nil
}

func New(diseasesRepository repositories.IDiseasesRepository) *Usecase {
	return &Usecase{diseasesRepository: diseasesRepository}
}
