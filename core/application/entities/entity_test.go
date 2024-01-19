package entities

import (
	valueobjects "health/core/application/value-objects"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type EntityStub struct {
	Entity
	name string
}

func TestIsGeneratingCreatedAt(t *testing.T) {

	now := time.Now()
	auditProps := AuditProps{
		CreatedAt: &now,
	}

	entity := NewEntity(&valueobjects.UniqueEntityUUID{}, &auditProps)
	assert.Equal(t, now, entity.CreatedAt, "Should set a inserted created at")
}

func TestIsGeneratingUpdatedAt(t *testing.T) {

	now := time.Now()
	auditProps := AuditProps{
		CreatedAt: &now,
		UpdatedAt: &time.Time{},
	}

	entity := NewEntity(&valueobjects.UniqueEntityUUID{}, &auditProps)
	assert.NotNil(t, entity.UpdatedAt, "Should set a inserted updated at")
}

func TestIsGeneratingDeletedAt(t *testing.T) {

	now := time.Now()
	auditProps := AuditProps{
		CreatedAt: &now,
		DeletedAt: &time.Time{},
	}

	entity := NewEntity(&valueobjects.UniqueEntityUUID{}, &auditProps)
	assert.NotNil(t, entity.DeletedAt, "Should set a inserted deleted at")
}
