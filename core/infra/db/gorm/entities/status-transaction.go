package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StatusTransaction struct {
	Uuid uuid.UUID `gorm:"column:uuid;primaryKey"`

	Id int `gorm:"autoIncrement;primaryKey"`

	Status uint8

	IdTransaction int `gorm:"column:id_transaction;not null"`

	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:now()"`
	UpdatedAt *time.Time     `gorm:"nullable;type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Transaction Transaction `gorm:"foreignKey:IdTransaction"`
}

func (u *StatusTransaction) TableName() string {
	return "accounting.status_transactions"
}
