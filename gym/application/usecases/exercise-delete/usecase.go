package exercisedelete

import (
	"health/gym/domain/repositories"
)

type Usecase struct {
	exercisesRepository repositories.IExercisesRepository
}

func (usecase *Usecase) Execute(id string) error {

	exercise, err := usecase.exercisesRepository.FindByID(id, nil)

	if err != nil {
		return err
	}

	return usecase.exercisesRepository.Delete(*exercise, nil)

}

func New(exercisesRepository repositories.IExercisesRepository) *Usecase {
	return &Usecase{exercisesRepository: exercisesRepository}
}
