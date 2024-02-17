package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateValidationCodeEntity(t *testing.T) {
	entity, err := NewValidationCode(nil, nil, nil, nil)
	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, entity.Entity.GetID(), "Should contain a valid id")

}

func TestSetValidationCodeCode(t *testing.T) {

	val := "aa"
	_, err := NewValidationCode(&val, nil, nil, nil)
	assert.NotNil(t, err, "Name should contain 6 chars")

	val = "aaaaaa"
	pacient, err := NewValidationCode(&val, nil, nil, nil)

	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, pacient.GetCode(), "entity should contain name")
}

func TestSetValidationCodeExpirationDate(t *testing.T) {

	val := ""
	validationCode, err := NewValidationCode(nil, &val, nil, nil)
	assert.NotNil(t, err, "Name should contain 11 chars")
	assert.Nil(t, validationCode, "Should not create entity")

	val = time.Now().Format(time.RFC3339)

	validationCode, err = NewValidationCode(nil, &val, nil, nil)

	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, validationCode.GetExpirationDate(), "entity should contain name")
}
