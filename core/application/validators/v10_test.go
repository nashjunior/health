package validators

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubValidator struct {
	Name string `validate:"required"`
	Age  int    `validate:"required,min=10`
}

func TestValidationWithErrors(t *testing.T) {
	stub := StubValidator{}

	validator := NewV10Validator(stub)

	validator.Validate()

	assert.NotNil(t, validator.Errors(), "Errors should not be empty")

}

func TestValidationWithoutErrors(t *testing.T) {
	stub := StubValidator{Name: "123", Age: 10}

	validator := NewV10Validator(stub)

	validator.Validate()

	assert.Nil(t, validator.Errors(), "Errors should not be empty")

}
