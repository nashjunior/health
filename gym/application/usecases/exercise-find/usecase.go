package exercisefind

import (
	"fmt"
	"health/gym/application/dtos"
	"health/gym/domain/repositories"
)

type Usecase struct {
	exercisesRepository repositories.IExercisesRepository
}

func (usecase *Usecase) Execute(id string) (*dtos.ExerciseOutput, error) {

	disease, err := usecase.exercisesRepository.FindByID(id, nil)

	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return &dtos.ExerciseOutput{
		Uuid:       disease.GetID().String(),
		Name:       disease.GetName(),
		CreatedAt:  disease.CreatedAt,
		UpdateddAt: disease.UpdatedAt,
		DeletedAt:  disease.DeletedAt,
	}, nil
}

func New(exercisesRepository repositories.IExercisesRepository) *Usecase {
	return &Usecase{exercisesRepository: exercisesRepository}
}
