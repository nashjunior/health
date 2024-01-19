package validators

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type V10 struct {
	schema any
	errors *[]ErrorField
}

func (v10Validator *V10) Validate() bool {
	validation := validator.New(validator.WithRequiredStructEnabled())

	err := validation.Struct(v10Validator.schema)

	if err != nil {
		errs := []ErrorField{}
		for _, err := range err.(validator.ValidationErrors) {
			errs = append(
				errs,
				ErrorField{Name: err.Field(), Error: err.Tag()},
			)
		}
		v10Validator.errors = &errs
		return false
	}

	return true
}

func (v10Validator *V10) Errors() *[]ErrorField {
	return v10Validator.errors
}

func NewV10Validator(structure any) IValidator {
	valueType := reflect.ValueOf(structure)

	if valueType.Kind() != reflect.Struct {
		return nil
	}

	return &V10{schema: structure}
}
