package suplementdelete

import (
	"health/nutrition/domain/repositories"
)

type Usecase struct {
	suplementsRepository repositories.ISuplementsRepository
}

func (usecase *Usecase) Execute(id string) error {

	disease, err := usecase.suplementsRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	return usecase.suplementsRepository.Delete(*disease, nil)

}

func New(suplementsRepository repositories.ISuplementsRepository) *Usecase {
	return &Usecase{suplementsRepository: suplementsRepository}
}
