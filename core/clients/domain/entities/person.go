package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
)

var validationPerson *validator.Validate

type Person struct {
	entities.Entity

	cpf    *string
	gender *string

	user      *User
	companies *[]Shareholder
}

func (person *Person) GetUser() User {
	return *person.user
}

func (person *Person) setCPF(cpf string) error {
	err := validationPerson.Var(cpf, "required,len=11")

	if err != nil {
		return err
	}

	person.cpf = &cpf

	return nil
}

func (person *Person) GetCPF() *string {
	return person.cpf
}

func (person *Person) setGender(gender string) error {
	err := validationPerson.Var(gender, "required,len=3")

	if err != nil {
		return err
	}

	person.gender = &gender

	return nil
}

func (person *Person) GetGender() *string {
	return person.gender
}

func (person *Person) Update(cpf *string, gender *string) error {
	var err error

	if cpf != nil {
		err = person.setCPF(*cpf)
		if err != nil {
			return err
		}

	}

	if gender != nil {
		err = person.setGender(*gender)
		if err != nil {
			return err
		}
	}

	now := time.Now()
	person.UpdatedAt = &now
	return nil

}

func NewPerson(cpf *string, gender *string, user *User, id *valueobjects.UniqueEntityUUID) (*Person, error) {
	validationPerson = validator.New(validator.WithRequiredStructEnabled())
	var err error

	person := &Person{}

	if cpf != nil {
		err = person.setCPF(*cpf)
		if err != nil {
			return nil, err
		}

	}

	if gender != nil {
		err = person.setGender(*gender)
		if err != nil {
			return nil, err
		}
	}

	person.Entity = entities.NewEntity(id, nil)
	person.user = user

	return person, nil
}
