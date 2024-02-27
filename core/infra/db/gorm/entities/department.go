package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	Uuid uuid.UUID `gorm:"column:uuid;primaryKey"`

	Id int `gorm:"autoIncrement;primaryKey"`

	Name string

	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:now()"`
	UpdatedAt *time.Time     `gorm:"nullable;type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (u *Department) TableName() string {
	return "public.departments"
}
