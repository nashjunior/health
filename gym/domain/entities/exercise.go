package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationExercise *validator.Validate

type Exercise struct {
	entities.Entity

	name *string
}

func (exercise *Exercise) SetName(name string) error {

	err := validationExercise.Var(name, "required")

	if err != nil {
		return err
	}

	exercise.name = &name

	return nil
}

func (exercise *Exercise) GetName() string {
	return *exercise.name
}

func (exercise *Exercise) Update(name *string) error {
	return exercise.SetName(*name)
}

func NewExercise(
	name *string,
	id *valueobjects.UniqueEntityUUID,
	time *entities.AuditProps,
) (*Exercise, error) {
	validationExercise = validator.New(validator.WithRequiredStructEnabled())
	exercise := &Exercise{}

	err := exercise.SetName(*name)

	if err != nil {
		exercise = nil
	} else {
		exercise.Entity = entities.NewEntity(id, time)
	}

	return exercise, err
}
