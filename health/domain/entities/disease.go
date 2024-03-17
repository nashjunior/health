package entities

import (
	"fmt"
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationDisease *validator.Validate

type Disease struct {
	entities.Entity

	name *string
}

func (disease *Disease) GetName() string {
	return *disease.name
}

func (disease *Disease) SetName(name string) error {
	fmt.Println(name)
	err := validationDisease.Var(name, "required")

	if err != nil {
		return err
	}

	disease.name = &name

	return nil
}

func (disease *Disease) Update(name *string) error {
	return disease.SetName(*name)
}

func NewDisease(
	name *string,
	id *valueobjects.UniqueEntityUUID,
	time *entities.AuditProps,
) (*Disease, error) {
	validationDisease = validator.New(validator.WithRequiredStructEnabled())
	disease := &Disease{}

	err := disease.SetName(*name)

	if err != nil {
		disease = nil
	} else {
		disease.Entity = entities.NewEntity(id, time)
	}

	return disease, err
}
