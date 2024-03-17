package createdisease

import (
	"health/health/application/dtos"
	"health/health/domain/entities"
	"health/health/domain/repositories"
)

type Input struct {
	Name *string
}

type Usecase struct {
	diseasesRepository repositories.IDiseasesRepository
}

func (usecase *Usecase) Execute(input Input) (*dtos.DiseaseOutput, error) {
	disease, err := entities.NewDisease(input.Name, nil, nil)

	if err != nil {
		return nil, err
	}

	err = usecase.diseasesRepository.Create(disease, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.DiseaseOutput{
		Uuid:      disease.GetID().String(),
		Name:      disease.GetName(),
		CreatedAt: disease.CreatedAt,
	}, nil
}

func New(diseasesRepository repositories.IDiseasesRepository) *Usecase {
	return &Usecase{diseasesRepository: diseasesRepository}
}
