package entities

import (
	"fmt"
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
)

type Status uint8

var validationStatusStatusTransaction *validator.Validate

const (
	Pending  Status = 0
	Rejected Status = 1
	Approved Status = 2
)

type StatusTransaction struct {
	entities.Entity
	id *int

	status *Status

	transaction *Transaction
}

func (transaction *StatusTransaction) SetInternalId(id int) {
	transaction.id = &id
}

func (transaction *StatusTransaction) GetInternalId() int {
	return *transaction.id
}

func (statusTransaction *StatusTransaction) GetStatus() Status {
	return *statusTransaction.status
}

func (statusTransaction *StatusTransaction) setStatus(status Status) error {
	switch status {
	case Pending, Rejected, Approved:
		{

			statusTransaction.status = &status
			break
		}
	default:
		return fmt.Errorf("Unrecognized status")
	}

	return nil
}

func (statusTransaction *StatusTransaction) GetTransaction() Transaction {
	return *statusTransaction.transaction
}

func (statusTransaction *StatusTransaction) Update(
	status *Status,
	transaction *Transaction,
) error {

	if status != nil {
		err := statusTransaction.setStatus(*status)
		if err != nil {
			return err
		}
		statusTransaction.status = status
	}

	if transaction != nil {
		statusTransaction.transaction = transaction
	}

	now := time.Now()
	statusTransaction.UpdatedAt = &now
	return nil
}

func NewStatusTransaction(
	status *Status,
	transaction *Transaction,
	id *valueobjects.UniqueEntityUUID,
) (*StatusTransaction, error) {
	validationStatusStatusTransaction = validator.New(validator.WithRequiredStructEnabled())

	statusTransaction := &StatusTransaction{}

	if status != nil {
		err := statusTransaction.setStatus(*status)
		if err != nil {
			return nil, err
		}
		statusTransaction.status = status
	}

	statusTransaction.transaction = transaction

	statusTransaction.Entity = entities.NewEntity(id, nil)

	return statusTransaction, nil
}
