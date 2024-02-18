package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	Uuid uuid.UUID `gorm:"column:uuid;primaryKey"`

	Id int `gorm:"autoIncrement;primaryKey"`

	Date  time.Time `gorm:"not null"`
	Value float64   `gorm:"not null"`

	IdTypeTransaction int `gorm:"column:id_type_transaction;not null"`

	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:now()"`
	UpdatedAt *time.Time     `gorm:"nullable;type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	TypeTransaction TypeTransaction `gorm:"foreignKey:IdTypeTransaction"`
}

func (u *Transaction) TableName() string {
	return "accounting.transactions"
}
