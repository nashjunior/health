package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationEquipament *validator.Validate

type Equipament struct {
	entities.Entity

	name *string
}

func (equipament *Equipament) SetName(name string) error {

	err := validationEquipament.Var(name, "required")

	if err != nil {
		return err
	}

	equipament.name = &name

	return nil
}

func (equipament *Equipament) GetName() string {
	return *equipament.name
}

func (equipament *Equipament) Update(name *string) error {
	return equipament.SetName(*name)
}

func NewEquipament(
	name *string,
	id *valueobjects.UniqueEntityUUID,
	time *entities.AuditProps,
) (*Equipament, error) {
	validationEquipament = validator.New(validator.WithRequiredStructEnabled())
	equipament := &Equipament{}

	err := equipament.SetName(*name)

	if err != nil {
		equipament = nil
	} else {
		equipament.Entity = entities.NewEntity(id, time)
	}

	return equipament, err
}
