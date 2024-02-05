package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreatePersonEntity(t *testing.T) {
	entity, err := NewPerson(nil, nil, nil, nil)
	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, entity.Entity.GetID(), "Should contain a valid id")

}

func TestSetPersonCPF(t *testing.T) {

	val := "aa"
	_, err := NewPerson(&val, nil, nil, nil)
	assert.NotNil(t, err, "Name should contain 11 chars")

	val = "aaaaaaaaaaa"
	pacient, err := NewPerson(&val, nil, nil, nil)

	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, pacient.GetCPF(), "entity should contain name")
}

func TestSetPersonGender(t *testing.T) {

	val := "aa"
	_, err := NewPerson(nil, &val, nil, nil)
	assert.NotNil(t, err, "Name should contain 11 chars")

	val = "aaa"
	pacient, err := NewPerson(nil, &val, nil, nil)

	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, pacient.GetGender(), "entity should contain name")
}

func TestUpdatePersonCpf(t *testing.T) {

	val := "aaaaaaaaaaa"
	person, _ := NewPerson(&val, nil, nil, nil)
	val = "bbbbbbbbbbb"

	err := person.Update(&val, nil)

	cpf := person.GetCPF()

	assert.Nil(t, err, "Name should contain required error")
	assert.Equal(t, val, *cpf, "entity should contain name")
}

func TestUpdatePersonGender(t *testing.T) {

	val := "aaa"
	person, _ := NewPerson(nil, &val, nil, nil)
	val = "bbb"

	err := person.Update(nil, &val)

	gender := person.GetGender()

	assert.Nil(t, err, "Name should contain required error")
	assert.Equal(t, val, *gender, "entity should contain name")
}
