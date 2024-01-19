package valueobjects

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIsUniqueUuidGeneratedWhenNull(t *testing.T) {
	instance1 := NewUniqueUUID(uuid.NullUUID{})

	assert.NotNil(t, instance1.Id, "Id should not be empty")
	_, error := uuid.Parse(instance1.Id.String())

	assert.Nil(t, error, "Id should be a valid uuid")
}

func TestIsUniqueUuidGeneratedWhenNotNull(t *testing.T) {
	uuidGen := uuid.NullUUID{UUID: uuid.New()}
	instance1 := NewUniqueUUID(uuidGen)

	assert.NotNil(t, instance1.Id, "Id should not be empty")
	_, error := uuid.Parse(instance1.Id.String())

	assert.Nil(t, error, "Id should be a valid uuid")

	assert.Equal(t, instance1.Id.String(), uuidGen.UUID.String(), "Should contain same uuid")
}

func TestIsUniqueUuidIsSameStruct(t *testing.T) {
	uuidGen := uuid.NullUUID{UUID: uuid.New()}
	instance1 := NewUniqueUUID(uuidGen)

	id1 := NewUniqueUUID(uuid.NullUUID{UUID: uuid.New()})
	isDifferent := instance1.Equals(&id1)

	assert.False(t, isDifferent, "Different instance")

	id2 := NewUniqueUUID(uuidGen)
	isEqual := instance1.Equals(&id2)
	assert.True(t, isEqual, "Should be equal instance")

}
