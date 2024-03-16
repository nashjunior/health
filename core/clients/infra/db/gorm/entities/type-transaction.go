package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TypeTransaction struct {
	Uuid uuid.UUID `gorm:"column:uuid;primaryKey"`

	Id int `gorm:"autoIncrement;primaryKey"`

	Name          string `gorm:"not null"`
	AccountType   uint8  `gorm:"not null;column:account_type"`
	OperationType uint8  `gorm:"not null;column:operation_type"`

	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:now()"`
	UpdatedAt *time.Time     `gorm:"nullable;type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *TypeTransaction) TableName() string {
	return "accounting.types_transactions"
}
