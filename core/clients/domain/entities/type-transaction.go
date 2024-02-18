package entities

import (
	"fmt"
	"health/core/application/entities"
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/go-playground/validator/v10"
)

type AccountType uint8
type OperationType uint8

const (
	Active   AccountType = 1
	Passive  AccountType = 2
	NetWorth AccountType = 3
	Revenue  AccountType = 4
)

const (
	Positive OperationType = 1
	Negative OperationType = 2
)

var validationTypeTransaction *validator.Validate

type TypeTransaction struct {
	entities.Entity

	id *int

	name          *string
	accountType   *AccountType
	operationType *OperationType
}

func (transactionType *TypeTransaction) SetInternalId(id int) {
	transactionType.id = &id
}

func (transactionType *TypeTransaction) GetInternalId() int {
	return *transactionType.id
}

func (transactionType *TypeTransaction) GetName() string {
	return *transactionType.name
}

func (transactionType *TypeTransaction) setName(name string) error {
	err := validationTypeTransaction.Var(name, "required")

	if err != nil {
		return err
	}

	transactionType.name = &name

	return nil
}

func (transactionType *TypeTransaction) setAccountType(accountType AccountType) error {
	switch accountType {
	case Active, Passive, NetWorth, Revenue:
		transactionType.accountType = &accountType
		break
	default:
		return fmt.Errorf("invalid enum value: %d", accountType)
	}

	return nil

}

func (transactionType *TypeTransaction) GetAccountType() AccountType {
	return *transactionType.accountType
}

func (transactionType *TypeTransaction) setOperationType(operationType OperationType) error {
	switch operationType {
	case Negative, Positive:
		transactionType.operationType = &operationType
		break
	default:
		return fmt.Errorf("invalid enum value: %d", operationType)
	}

	return nil
}

func (transactionType *TypeTransaction) GetOperationType() OperationType {
	return *transactionType.operationType
}

func (transactionType *TypeTransaction) Update(
	name *string,
	accountType *AccountType,
	operationType *OperationType,
) error {

	if name != nil {
		err := transactionType.setName(*name)
		if err != nil {
			return err
		}
	}

	if accountType != nil {
		err := transactionType.setAccountType(*accountType)
		if err != nil {
			return err
		}
	}

	now := time.Now()
	transactionType.UpdatedAt = &now
	return nil
}

func NewTypeTransaction(
	name *string,
	accountType *AccountType,
	operationType *OperationType,
	id *valueobjects.UniqueEntityUUID,
) (*TypeTransaction, error) {

	validationTypeTransaction = validator.New(validator.WithRequiredStructEnabled())

	typeTransaction := &TypeTransaction{}

	if name != nil {
		err := typeTransaction.setName(*name)
		if err != nil {
			return nil, err
		}
	}

	if operationType != nil {
		err := typeTransaction.setOperationType(*operationType)
		if err != nil {
			return nil, err
		}
	}

	if accountType != nil {
		err := typeTransaction.setAccountType(*accountType)
		if err != nil {
			return nil, err
		}
	}

	typeTransaction.Entity = entities.NewEntity(id, nil)

	return typeTransaction, nil
}
