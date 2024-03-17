package inmemory

import (
	"health/core/clients/domain/entities"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var repoUsers *UsersInMemoryRepository

func setupTestUsers() {
	repoUsers = NewUsersInMemoryRepository()
}

func TestShouldCreateUserInstance(t *testing.T) {
	setupTestUsers()

	cpf := "00000000000"
	entity, _ := entities.NewUser(&cpf, nil)

	err := repoUsers.Create(*entity, nil)

	assert.Nil(t, err, "Should not have error")
	assert.Len(t, repoUsers.Items, 1, "Repository size should be 1")
	assert.Equal(t, *entity, repoUsers.Items[0], "Should be same entity")
}

func TestShouldFindByUUIDUserInstance(t *testing.T) {
	setupTestUsers()

	cpf := "00000000000"
	entity, _ := entities.NewUser(&cpf, nil)

	repoUsers.Create(*entity, nil)

	_, err := repoUsers.FindByUUID(uuid.New(), nil)

	assert.NotNil(t, err, "Should not found a entity by UUID")

	schedule, err := repoUsers.FindByUUID(entity.GetID(), nil)
	assert.Nil(t, err, "Should found a entity by UUID")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldFindByIDUserInstance(t *testing.T) {
	setupTestUsers()

	cpf := "00000000000"
	entity, _ := entities.NewUser(&cpf, nil)

	repoUsers.Create(*entity, nil)

	_, err := repoUsers.FindByID("", nil)

	assert.NotNil(t, err, "Should not found a entity by string")

	schedule, err := repoUsers.FindByID(entity.GetID().String(), nil)
	assert.Nil(t, err, "Should found a entity by string")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldUpdateUserInstance(t *testing.T) {
	setupTestUsers()

	cpf := "00000000000"
	entity, _ := entities.NewUser(&cpf, nil)

	repoUsers.Create(*entity, nil)

	cpf = "11111111111"
	entity.Update(&cpf)

	err := repoUsers.Update(*entity, nil)

	assert.Nil(t, err, "Should throw no error after update entity")
	assert.Equal(t, *entity, repoUsers.Items[0], "They are not the same")
}

func TestShouldDeleteUserInstance(t *testing.T) {
	setupTestUsers()

	cpf := "00000000000"
	entity, _ := entities.NewUser(&cpf, nil)

	repoUsers.Create(*entity, nil)

	err := repoUsers.Delete(*entity, nil)

	assert.Nil(t, err, "Should not throw an error on delete entity")
	_, err = repoUsers.FindByUUID(entity.GetID(), nil)

	assert.NotNil(t, err, "Should throw an error on find entity")
}
