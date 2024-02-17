package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
)

var validationCodeVars *validator.Validate

type ValidationCode struct {
	entities.Entity

	code           *string
	expirationDate *time.Time
	active         bool

	user *User
}

func (validationCode *ValidationCode) GetPerson() User {
	return *validationCode.user
}

func (validationCode *ValidationCode) setCode(code string) error {
	err := validationCodeVars.Var(code, "required,len=6")

	if err != nil {
		return err
	}

	validationCode.code = &code

	return nil
}
func (validationCode *ValidationCode) GetCode() string { return *validationCode.code }

func (validationCode *ValidationCode) setExpirationDate(expirationDate string) error {

	err := validationCodeVars.Var(expirationDate, "required,datetime="+time.RFC3339)

	if err != nil {
		return err
	}

	date, _ := time.Parse(time.RFC3339, expirationDate)

	validationCode.expirationDate = &date

	return nil

}
func (validationCode *ValidationCode) GetExpirationDate() time.Time {
	return *validationCode.expirationDate
}

func (validationCode *ValidationCode) activate() {
	validationCode.active = true
}

func (validationCode *ValidationCode) deactivate() {
	validationCode.active = false
}

func NewValidationCode(code *string, expirationDate *string, user *User, id *valueobjects.UniqueEntityUUID) (*ValidationCode, error) {
	validationCodeVars = validator.New(validator.WithRequiredStructEnabled())
	var err error

	validationCode := &ValidationCode{}

	if code != nil {
		err = validationCode.setCode(*code)
		if err != nil {
			return nil, err
		}
	}

	if expirationDate != nil {
		err = validationCode.setExpirationDate(*expirationDate)
		if err != nil {
			return nil, err
		}
	}

	validationCode.Entity = entities.NewEntity(id, nil)
	validationCode.user = user
	validationCode.deactivate()

	return validationCode, nil
}
