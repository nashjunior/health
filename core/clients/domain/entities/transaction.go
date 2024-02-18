package entities

import (
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
)

var validationTransaction *validator.Validate

type Transaction struct {
	entities.Entity

	id *int

	date  *time.Time
	value *float64

	transactionType *TypeTransaction
}

func (transaction *Transaction) SetInternalId(id int) {
	transaction.id = &id
}

func (transaction *Transaction) GetInternalId() int {
	return *transaction.id
}

func (transaction *Transaction) GetDate() time.Time {
	return *transaction.date
}

func (transaction *Transaction) setDate(date time.Time) error {
	err := validationTransaction.Var(date.Format(time.RFC3339), "required,datetime="+time.RFC3339)

	if err != nil {
		return err
	}

	transaction.date = &date
	return nil
}

func (transaction *Transaction) GetValue() float64 {
	return *transaction.value
}

func (transaction *Transaction) setValue(value float64) error {
	err := validationTransaction.Var(value, "required")

	if err != nil {
		return err
	}

	transaction.value = &value
	return nil
}

func (transaction *Transaction) GetTypeTransaction() TypeTransaction {
	return *transaction.transactionType
}

func (transaction *Transaction) Update(
	date *time.Time,
	value *float64,
	transactionType *TypeTransaction,
) error {

	if date != nil {
		err := transaction.setDate(*date)
		if err != nil {
			return err
		}
	}

	if value != nil {
		err := transaction.setValue(*value)
		if err != nil {
			return err
		}
	}

	if transactionType != nil {
		transaction.transactionType = transactionType
	}

	now := time.Now()
	transaction.UpdatedAt = &now
	return nil
}

func NewTransaction(
	date *time.Time,
	value *float64,
	transactionType *TypeTransaction,
	id *valueobjects.UniqueEntityUUID,
) (*Transaction, error) {
	validationTransaction = validator.New(validator.WithRequiredStructEnabled())

	transaction := &Transaction{}

	if date != nil {
		err := transaction.setDate(*date)
		if err != nil {
			return nil, err
		}
	}

	if value != nil {
		err := transaction.setValue(*value)
		if err != nil {
			return nil, err
		}
	}

	transaction.transactionType = transactionType
	transaction.Entity = entities.NewEntity(id, nil)

	return transaction, nil
}
