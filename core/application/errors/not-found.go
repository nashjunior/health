package errors

import (
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Message    string
	StatusCode uint16
}

func (r *NotFoundError) Error() string {

	return fmt.Sprintf("Validation error")
}

func NewNotFoundError(message string) error {
	return &NotFoundError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}
