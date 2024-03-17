package exercisecreate

import (
	"errors"
	"health/gym/application/dtos"
	"health/gym/domain/entities"
	"health/gym/domain/repositories"
)

type Input struct {
	Name *string
}

type Usecase struct {
	exercisesRepository repositories.IExercisesRepository
}

func (usecase *Usecase) Execute(input Input) (*dtos.ExerciseOutput, error) {

	previousExercise, _ := usecase.exercisesRepository.FindByName(*input.Name, nil)

	if previousExercise != nil {
		return nil, errors.New("A equipament with this name already exists")
	}

	exercise, err := entities.NewExercise(input.Name, nil, nil)

	if err != nil {
		return nil, err
	}

	err = usecase.exercisesRepository.Create(exercise, nil)

	if err != nil {
		return nil, err
	}

	return &dtos.ExerciseOutput{
		Uuid:      exercise.GetID().String(),
		Name:      exercise.GetName(),
		CreatedAt: exercise.CreatedAt,
	}, nil
}

func New(exercisesRepository repositories.IExercisesRepository) *Usecase {
	return &Usecase{exercisesRepository: exercisesRepository}
}
