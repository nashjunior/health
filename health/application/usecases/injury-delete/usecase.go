package injurydelete

import (
	"health/health/domain/repositories"
)

type Usecase struct {
	injuriesRepository repositories.IInjuriesRepository
}

func (usecase *Usecase) Execute(id string) error {

	disease, err := usecase.injuriesRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	return usecase.injuriesRepository.Delete(*disease, nil)

}

func New(injuriesRepository repositories.IInjuriesRepository) *Usecase {
	return &Usecase{injuriesRepository: injuriesRepository}
}
