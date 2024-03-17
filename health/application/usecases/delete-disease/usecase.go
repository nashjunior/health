package deletedisease

import (
	"health/health/domain/repositories"
)

type Usecase struct {
	diseasesRepository repositories.IDiseasesRepository
}

func (usecase *Usecase) Execute(id string) error {

	disease, err := usecase.diseasesRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	return usecase.diseasesRepository.Delete(*disease, nil)

}

func New(diseasesRepository repositories.IDiseasesRepository) *Usecase {
	return &Usecase{diseasesRepository: diseasesRepository}
}
