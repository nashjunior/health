package errors

import (
	"fmt"
	"health/core/application/validators"
	"net/http"
)

type ValidationError struct {
	Errors     []validators.ErrorField
	StatusCode uint16
}

func (r *ValidationError) Error() string {

	return fmt.Sprintf("Validation error")
}

func NewValidationError(errors []validators.ErrorField) error {
	return &ValidationError{
		StatusCode: http.StatusUnprocessableEntity,
		Errors:     errors,
	}
}
