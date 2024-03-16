package inmemory

import (
	"health/core/clients/domain/entities"
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var repoPersons *PersonsInMemoryRepository

func setupTestPersons() {
	repoPersons = NewPersonsInMemoryRepository()
}

func TestShouldCreatePersonInstance(t *testing.T) {
	setupTestPersons()

	cnpj := "00000000000"
	entity, _ := entities.NewPerson(&cnpj, nil, nil, nil)

	err := repoPersons.Create(*entity, nil)

	assert.Nil(t, err, "Should not have error")
	assert.Len(t, repoPersons.Items, 1, "Repository size should be 1")
	assert.Equal(t, *entity, repoPersons.Items[0], "Should be same entity")
}

func TestShouldFindByUUIDPersonInstance(t *testing.T) {
	setupTestPersons()

	cnpj := "00000000000"
	entity, _ := entities.NewPerson(&cnpj, nil, nil, nil)

	repoPersons.Create(*entity, nil)

	_, err := repoPersons.FindByUUID(uuid.New(), nil)

	assert.NotNil(t, err, "Should not found a entity by UUID")

	schedule, err := repoPersons.FindByUUID(entity.GetID(), nil)
	assert.Nil(t, err, "Should found a entity by UUID")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldFindByIDPersonInstance(t *testing.T) {
	setupTestPersons()

	cnpj := "00000000000"
	entity, _ := entities.NewPerson(&cnpj, nil, nil, nil)

	repoPersons.Create(*entity, nil)

	_, err := repoPersons.FindByID("", nil)

	assert.NotNil(t, err, "Should not found a entity by string")

	schedule, err := repoPersons.FindByID(entity.GetID().String(), nil)
	assert.Nil(t, err, "Should found a entity by string")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldDeletePersonInstance(t *testing.T) {
	setupTestPersons()

	cnpj := "00000000000"
	entity, _ := entities.NewPerson(&cnpj, nil, nil, nil)

	repoPersons.Create(*entity, nil)

	err := repoPersons.Delete(*entity, nil)

	assert.Nil(t, err, "Should not throw an error on delete entity")
	_, err = repoPersons.FindByUUID(entity.GetID(), nil)

	assert.NotNil(t, err, "Should throw an error on find entity")
}

func TestShouldDeleteManyPersonsInstance(t *testing.T) {
	setupTestPersons()

	var entitiesToCreate []entities.Person

	for i := 0; i < 10; i++ {
		cnpj := strings.Repeat(strconv.Itoa(i), 11)

		entity, _ := entities.NewPerson(&cnpj, nil, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoPersons.CreateMany(entitiesToCreate, nil)

	err := repoPersons.DeleteMany(entitiesToCreate, nil)

	assert.Nil(t, err, "should not throw an error")

	_, err = repoPersons.FindByUUID(entitiesToCreate[0].GetID(), nil)
	assert.NotNil(t, err, "should throw an error on find by id")
}

func TestShouldFindAllPersons(t *testing.T) {
	setupTestPersons()

	var entitiesToCreate []entities.Person

	for i := 0; i < 10; i++ {
		cnpj := strings.Repeat(strconv.Itoa(i), 11)

		entity, _ := entities.NewPerson(&cnpj, nil, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoPersons.CreateMany(entitiesToCreate, nil)

	items := repoPersons.Find(nil, nil)
	assert.Len(t, items, 10, "Should return no items if not found by startDate")

}
