package dtos

import "time"

type EquipamentOutput struct {
	Uuid string `json:"id"`
	Name string `json:"name"`

	CreatedAt  time.Time  `json:"created_at"`
	UpdateddAt *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
}
