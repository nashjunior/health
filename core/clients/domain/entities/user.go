package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var validationUser *validator.Validate

type User struct {
	entities.Entity

	name         *string
	birthdayDate *time.Time
}

func (user *User) GetName() string {
	return *user.name
}

func (user *User) setName(name string) error {
	err := validationUser.Var(name, "required,min=3")

	if err != nil {
		return err
	}

	user.name = &name

	return nil
}

func (user *User) GetBirthdayDate() *time.Time {
	return user.birthdayDate
}

func (user *User) Update(name *string) error {
	var err error
	if name != nil {
		err = user.setName(*name)
	}

	if err != nil {
		return err
	}

	now := time.Now()
	user.UpdatedAt = &now

	return nil
}

func NewUser(name *string, id *valueobjects.UniqueEntityUUID) (*User, error) {
	validationUser = validator.New(validator.WithRequiredStructEnabled())

	user := &User{}

	var err error

	if name != nil {
		err = user.setName(*name)
	}

	if err != nil {
		user = nil
		return user, err
	}

	if id == nil {
		newId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: uuid.New()})
		id = &newId

	}

	user.Entity = entities.NewEntity(id, nil)

	return user, nil
}
