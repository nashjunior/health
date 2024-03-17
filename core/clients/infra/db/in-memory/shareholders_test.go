package inmemory

import (
	"health/core/clients/domain/entities"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var repoShareHolders *ShareholdersInMemoryRepository

func setupTestShareholders() {
	repoShareHolders = NewShareholdersInMemoryRepository()
}

func TestShouldCreateShareholderInstance(t *testing.T) {
	setupTestShareholders()

	entity, _ := entities.NewShareHolder(nil, nil, nil)

	err := repoShareHolders.Create(*entity, nil)

	assert.Nil(t, err, "Should not have error")
	assert.Len(t, repoShareHolders.Items, 1, "Repository size should be 1")
	assert.Equal(t, *entity, repoShareHolders.Items[0], "Should be same entity")
}

func TestShouldFindByUUIDShareholderInstance(t *testing.T) {
	setupTestShareholders()

	entity, _ := entities.NewShareHolder(nil, nil, nil)

	repoShareHolders.Create(*entity, nil)

	_, err := repoShareHolders.FindByUUID(uuid.New(), nil)

	assert.NotNil(t, err, "Should not found a entity by UUID")

	schedule, err := repoShareHolders.FindByUUID(entity.GetID(), nil)
	assert.Nil(t, err, "Should found a entity by UUID")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldFindByIDShareholderInstance(t *testing.T) {
	setupTestShareholders()

	entity, _ := entities.NewShareHolder(nil, nil, nil)

	repoShareHolders.Create(*entity, nil)

	_, err := repoShareHolders.FindByID("", nil)

	assert.NotNil(t, err, "Should not found a entity by string")

	schedule, err := repoShareHolders.FindByID(entity.GetID().String(), nil)
	assert.Nil(t, err, "Should found a entity by string")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldUpdateShareholderInstance(t *testing.T) {
	setupTestShareholders()

	entity, _ := entities.NewShareHolder(nil, nil, nil)

	repoShareHolders.Create(*entity, nil)

	now := time.Now()
	entity.UpdatedAt = &now

	err := repoShareHolders.Update(*entity, nil)

	assert.Nil(t, err, "Should throw no error after update entity")
	assert.Equal(t, *entity, repoShareHolders.Items[0], "They are not the same")
}

func TestShouldDeleteShareholderInstance(t *testing.T) {
	setupTestShareholders()

	entity, _ := entities.NewShareHolder(nil, nil, nil)

	repoShareHolders.Create(*entity, nil)

	err := repoShareHolders.Delete(*entity, nil)

	assert.Nil(t, err, "Should not throw an error on delete entity")
	_, err = repoShareHolders.FindByUUID(entity.GetID(), nil)

	assert.NotNil(t, err, "Should throw an error on find entity")
}

func TestShouldDeleteManyShareholdersInstance(t *testing.T) {
	setupTestShareholders()

	var entitiesToCreate []entities.Shareholder

	for i := 0; i < 10; i++ {

		entity, _ := entities.NewShareHolder(nil, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoShareHolders.CreateMany(entitiesToCreate, nil)

	err := repoShareHolders.DeleteMany(entitiesToCreate, nil)

	assert.Nil(t, err, "should not throw an error")

	_, err = repoShareHolders.FindByUUID(entitiesToCreate[0].GetID(), nil)
	assert.NotNil(t, err, "should throw an error on find by id")
}

func TestShouldFindAllShareholders(t *testing.T) {
	setupTestShareholders()

	var entitiesToCreate []entities.Shareholder

	for i := 0; i < 10; i++ {

		entity, _ := entities.NewShareHolder(nil, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoShareHolders.CreateMany(entitiesToCreate, nil)

	items := repoShareHolders.Find(nil, nil)

	assert.Len(t, items, 10, "Should return no items if not found by startDate")

}

func TestShouldFindAllShareholdersByPerson(t *testing.T) {
	setupTestShareholders()

	cpf := "00000000000"
	gender := "MAS"
	person, _ := entities.NewPerson(&cpf, &gender, nil, nil)

	var entitiesToCreate []entities.Shareholder
	for i := 0; i < 10; i++ {

		entity, _ := entities.NewShareHolder(person, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoShareHolders.CreateMany(entitiesToCreate, nil)

	customPerson, _ := entities.NewPerson(nil, nil, nil, nil)
	items := repoShareHolders.FindAllByPerson(*customPerson, nil)
	assert.Len(t, items, 0, "should find no customer roles with non existent role associated")

	items = repoShareHolders.FindAllByPerson(*person, nil)
	assert.Len(t, items, 10, "Should return all items founded with customer role")
}

func TestShouldFindAllShareholdersByRole(t *testing.T) {
	setupTestShareholders()

	cpf := "00000000000000"
	enterprise, _ := entities.NewEnterprise(&cpf, nil, nil, nil)

	var entitiesToCreate []entities.Shareholder
	for i := 0; i < 10; i++ {
		entity, _ := entities.NewShareHolder(nil, enterprise, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoShareHolders.CreateMany(entitiesToCreate, nil)

	customEnterprise, _ := entities.NewEnterprise(nil, nil, nil, nil)
	items := repoShareHolders.FindAllByEnterprises(*customEnterprise, nil)
	assert.Len(t, items, 0, "should find no customer roles with non existent role associated")

	items = repoShareHolders.FindAllByEnterprises(*enterprise, nil)
	assert.Len(t, items, 10, "Should return all items founded with customer role")
}
