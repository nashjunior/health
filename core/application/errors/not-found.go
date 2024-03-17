package errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type NotFoundError struct {
	Message    string
	StatusCode uint16
}

func (r *NotFoundError) Error() string {
	a := map[string]any{
		"message": r.Message,
	}

	resp, err := json.Marshal(a)

	if err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%s", resp)
}

func NewNotFoundError(message string) error {
	return &NotFoundError{
		Message:    message,
		StatusCode: http.StatusNotFound,
	}
}
