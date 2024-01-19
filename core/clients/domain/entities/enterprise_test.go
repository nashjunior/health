package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldCreateEnterpriseEntity(t *testing.T) {
	entity, err := NewEnterprise(nil, nil, nil, nil)
	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, entity.Entity.GetID(), "Should contain a valid id")

}

func TestSetEnterpriseCNPJ(t *testing.T) {

	val := "aa"
	_, err := NewEnterprise(&val, nil, nil, nil)
	assert.NotNil(t, err, "Name should contain 11 chars")

	val = "aaaaaaaaaaaaaa"
	pacient, err := NewEnterprise(&val, nil, nil, nil)

	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, pacient.GetCnpj(), "entity should contain name")
}

func TestSetEnterpriseSocialReason(t *testing.T) {

	val := "aa"
	_, err := NewEnterprise(nil, &val, nil, nil)
	assert.NotNil(t, err, "Name should contain 11 chars")

	val = "000"
	pacient, err := NewEnterprise(nil, &val, nil, nil)

	assert.Nil(t, err, "Name should contain required error")
	assert.NotNil(t, pacient.GetSocialReason(), "entity should contain name")
}

func TestUpdateEnterpriseCnpj(t *testing.T) {

	val := "aaaaaaaaaaaaaa"

	Enterprise, _ := NewEnterprise(&val, nil, nil, nil)
	val = "bbbbbbbbbbbbbb"

	err := Enterprise.Update(&val, nil)

	cnpj := Enterprise.GetCnpj()

	assert.Nil(t, err, "Name should contain required error")
	assert.Equal(t, val, *cnpj, "entity should contain name")
}

func TestUpdateEnterpriseSocialReason(t *testing.T) {

	val := "aaa"
	Enterprise, _ := NewEnterprise(nil, &val, nil, nil)
	val = "bbb"

	err := Enterprise.Update(nil, &val)

	socialReason := Enterprise.GetSocialReason()

	assert.Nil(t, err, "Name should contain required error")
	assert.Equal(t, val, *socialReason, "entity should contain name")
}
