package equipamentdelete

import (
	"health/gym/domain/repositories"
)

type Usecase struct {
	equipamentsRepository repositories.IEquipamentsRepository
}

func (usecase *Usecase) Execute(id string) error {

	disease, err := usecase.equipamentsRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	return usecase.equipamentsRepository.Delete(*disease, nil)

}

func New(equipamentsRepository repositories.IEquipamentsRepository) *Usecase {
	return &Usecase{equipamentsRepository: equipamentsRepository}
}
