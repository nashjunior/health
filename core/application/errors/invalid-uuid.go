package errors

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type InvalidUUidError struct {
	Id         uuid.UUID
	StatusCode uint16
}

func (r *InvalidUUidError) Error() string {

	return fmt.Sprintf("Validation error")
}

func NewInvalidUUidError(id uuid.UUID) error {
	return &InvalidUUidError{
		Id:         id,
		StatusCode: http.StatusNotFound,
	}
}
