package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DepartmentHierarchy struct {
	Uuid uuid.UUID `gorm:"column:uuid;primaryKey"`

	Id int `gorm:"autoIncrement;primaryKey"`

	IdDepartment int  `gorm:"column:id_department"`
	IdManager    *int `gorm:"column:id_manager"`

	CreatedAt time.Time      `gorm:"type:timestamp;column:created_at;default:now()"`
	UpdatedAt *time.Time     `gorm:"nullable;type:timestamp;column:updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Department Department  `gorm:"foreignKey:IdDepartment"`
	Manager    *Department `gorm:"foreignKey:IdManager"`
}

func (u *DepartmentHierarchy) TableName() string {
	return "public.departments_closure"
}
