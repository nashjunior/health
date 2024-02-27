package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"

	"github.com/go-playground/validator/v10"
)

var validationDepartment *validator.Validate

type Department struct {
	entities.Entity
	id *int

	name *string

	managers     *[]Department
	subordinates *[]Department
}

func (department *Department) SetManagers(managers []Department) {
	department.managers = &managers
}

func (department *Department) GetManagers() []Department {
	return *department.managers
}

func (department *Department) SetSubordinates(subordinates []Department) {
	department.subordinates = &subordinates
}

func (department *Department) GetSubordinates() []Department {
	return *department.subordinates
}

func (department *Department) SetInternalId(id int) {
	department.id = &id
}

func (department *Department) GetInternalId() int {
	return *department.id
}

func (department *Department) setName(name string) error {
	err := validationDepartment.Var(name, "required")

	if err != nil {
		return err
	}

	department.name = &name
	return nil
}

func (department *Department) GetName() string {
	return *department.name
}

func (department *Department) Update(name *string) error {
	if name != nil {
		err := department.setName(*name)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewDepartment(
	name *string,
	managers *[]Department,
	subordinates *[]Department,
	id *valueobjects.UniqueEntityUUID,
) (*Department, error) {
	department := &Department{}
	validationDepartment = validator.New()

	if name != nil {
		err := department.setName(*name)
		if err != nil {
			return nil, err
		}
	}

	department.managers = managers
	department.subordinates = subordinates

	department.Entity = entities.NewEntity(id, nil)
	return department, nil
}
