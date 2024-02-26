package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobHierarchy struct {
	Uuid uuid.UUID `gorm:"column:uuid;primaryKey"`

	Id int `gorm:"autoIncrement;primaryKey"`

	IdJob     int  `gorm:"column:id_job"`
	IdManager *int `gorm:"column:id_manager"`

	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:now()"`
	UpdatedAt *time.Time     `gorm:"nullable;type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Job     Job  `gorm:"foreignKey:IdJob"`
	Manager *Job `gorm:"foreignKey:IdManager"`
}

func (u *JobHierarchy) TableName() string {
	return "public.jobs_closure"
}
