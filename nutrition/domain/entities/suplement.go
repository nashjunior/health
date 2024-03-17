package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationSuplement *validator.Validate

type Suplement struct {
	entities.Entity

	name *string
}

func (suplement *Suplement) SetName(name string) error {

	err := validationSuplement.Var(name, "required")

	if err != nil {
		return err
	}

	suplement.name = &name

	return nil
}

func (suplement *Suplement) GetName() string {
	return *suplement.name
}

func (suplement *Suplement) Update(name *string) error {
	return suplement.SetName(*name)
}

func NewSuplement(
	name *string,
	id *valueobjects.UniqueEntityUUID,
	time *entities.AuditProps,
) (*Suplement, error) {
	validationSuplement = validator.New(validator.WithRequiredStructEnabled())
	suplement := &Suplement{}

	err := suplement.SetName(*name)

	if err != nil {
		suplement = nil
	} else {
		suplement.Entity = entities.NewEntity(id, time)
	}

	return suplement, err
}
