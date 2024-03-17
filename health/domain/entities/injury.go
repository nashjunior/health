package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationInjury *validator.Validate

type Injury struct {
	entities.Entity

	name *string
}

func (injury *Injury) GetName() string {
	return *injury.name
}

func (injury *Injury) SetName(name string) error {

	err := validationInjury.Var(name, "required")

	if err != nil {
		return err
	}

	injury.name = &name

	return nil
}

func (injury *Injury) Update(name *string) error {
	return injury.SetName(*name)
}

func NewInjury(
	name *string,
	id *valueobjects.UniqueEntityUUID,
	time *entities.AuditProps,
) (*Injury, error) {
	validationInjury = validator.New(validator.WithRequiredStructEnabled())
	injury := &Injury{}

	err := injury.SetName(*name)

	if err != nil {
		injury = nil
	} else {
		injury.Entity = entities.NewEntity(id, time)
	}

	return injury, err
}
