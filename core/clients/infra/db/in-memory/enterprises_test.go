package inmemory

import (
	"health/core/clients/domain/entities"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var repoEnterprises *EnterprisesInMemoryRepository

func setupTestEnterprises() {
	repoEnterprises = NewEnterprisesInMemoryRepository()
}

func TestShouldCreateEnterpriseInstance(t *testing.T) {
	setupTestEnterprises()

	socialReason := "aaa"
	cnpj := "00000000000000"
	entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)

	err := repoEnterprises.Create(*entity, nil)

	assert.Nil(t, err, "Should not have error")
	assert.Len(t, repoEnterprises.Items, 1, "Repository size should be 1")
	assert.Equal(t, *entity, repoEnterprises.Items[0], "Should be same entity")
}

func TestShouldFindByUUIDEnterpriseInstance(t *testing.T) {
	setupTestEnterprises()

	socialReason := "aaa"
	cnpj := "00000000000000"
	entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)

	repoEnterprises.Create(*entity, nil)

	_, err := repoEnterprises.FindByUUID(uuid.New(), nil)

	assert.NotNil(t, err, "Should not found a entity by UUID")

	schedule, err := repoEnterprises.FindByUUID(entity.GetID(), nil)
	assert.Nil(t, err, "Should found a entity by UUID")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldFindByIDEnterpriseInstance(t *testing.T) {
	setupTestEnterprises()

	socialReason := "aaa"
	cnpj := "00000000000000"
	entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)

	repoEnterprises.Create(*entity, nil)

	_, err := repoEnterprises.FindByID("", nil)

	assert.NotNil(t, err, "Should not found a entity by string")

	schedule, err := repoEnterprises.FindByID(entity.GetID().String(), nil)
	assert.Nil(t, err, "Should found a entity by string")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldUpdateEnterpriseInstance(t *testing.T) {
	setupTestEnterprises()

	socialReason := "aaa"
	cnpj := "00000000000000"
	entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)

	repoEnterprises.Create(*entity, nil)

	cnpj = "11111111111111"
	entity.Update(&cnpj, nil)

	err := repoEnterprises.Update(*entity, nil)

	assert.Nil(t, err, "Should throw no error after update entity")
	assert.Equal(t, *entity, repoEnterprises.Items[0], "They are not the same")
}

func TestShouldDeleteEnterpriseInstance(t *testing.T) {
	setupTestEnterprises()

	socialReason := "aaa"
	cnpj := "00000000000000"
	entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)

	repoEnterprises.Create(*entity, nil)

	err := repoEnterprises.Delete(*entity, nil)

	assert.Nil(t, err, "Should not throw an error on delete entity")
	_, err = repoEnterprises.FindByUUID(entity.GetID(), nil)

	assert.NotNil(t, err, "Should throw an error on find entity")
}

func TestShouldDeleteManyEnterprisesInstance(t *testing.T) {
	setupTestEnterprises()

	var entitiesToCreate []entities.Enterprise

	for i := 0; i < 10; i++ {
		cnpj := strings.Repeat(strconv.Itoa(i), 14)
		socialReason := "aaa"

		entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoEnterprises.CreateMany(entitiesToCreate, nil)

	err := repoEnterprises.DeleteMany(entitiesToCreate, nil)

	assert.Nil(t, err, "should not throw an error")

	_, err = repoEnterprises.FindByUUID(entitiesToCreate[0].GetID(), nil)
	assert.NotNil(t, err, "should throw an error on find by id")
}

func TestShouldFindAllEnterprises(t *testing.T) {
	setupTestEnterprises()

	var entitiesToCreate []entities.Enterprise

	for i := 0; i < 10; i++ {
		cnpj := strings.Repeat(strconv.Itoa(i), 14)
		socialReason := "aaa"

		entity, _ := entities.NewEnterprise(&cnpj, &socialReason, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoEnterprises.CreateMany(entitiesToCreate, nil)

	items := repoEnterprises.Find(nil, nil)

	assert.Len(t, items, 10, "Should return no items if not found by startDate")

}
