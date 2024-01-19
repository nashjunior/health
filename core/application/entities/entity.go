package entities

import (
	valueobjects "health/core/application/value-objects"
	"time"

	"github.com/google/uuid"
)

type AuditProps struct {
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

type Entity struct {
	UniqueEntityUUID valueobjects.UniqueEntityUUID
	CreatedAt        time.Time  `json:"created_at"`
	UpdatedAt        *time.Time `json:"updated_at"`
	DeletedAt        *time.Time `json:"deleted_at"`
}

func (entity *Entity) GetID() uuid.UUID {
	return entity.UniqueEntityUUID.Id
}

func NewEntity(
	id *valueobjects.UniqueEntityUUID,
	audit *AuditProps,
) Entity {

	var uniqueId valueobjects.UniqueEntityUUID
	if id == nil {
		uniqueId = valueobjects.NewUniqueUUID(uuid.NullUUID{
			UUID: uuid.New(),
		})
	} else {
		uniqueId = *id
	}

	var createdAt time.Time
	var updatedAt *time.Time
	var deletedAt *time.Time

	if audit != nil {
		if audit.CreatedAt != nil {
			createdAt = *audit.CreatedAt
		}

		updatedAt = audit.UpdatedAt
		deletedAt = audit.DeletedAt
	} else {
		createdAt = time.Now()
	}

	return Entity{
		UniqueEntityUUID: uniqueId,
		CreatedAt:        createdAt,
		UpdatedAt:        updatedAt,
		DeletedAt:        deletedAt,
	}
}
