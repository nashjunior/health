package inmemory

import (
	ent "health/core/application/entities"
	"health/core/clients/domain/entities"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var repoValidationsCodes *ValidationsCodesInMemoryRepository

func setupTestValidationsCodes() {
	repoValidationsCodes = NewValidationsCodesInMemoryRepository()
}

func TestShouldCreateValidationCodeInstance(t *testing.T) {
	setupTestValidationsCodes()

	code := "000000"
	now := time.Now().Format(time.RFC3339)
	entity, _ := entities.NewValidationCode(&code, &now, nil, nil)

	err := repoValidationsCodes.Create(*entity, nil)

	assert.Nil(t, err, "Should not have error")
	assert.Len(t, repoValidationsCodes.Items, 1, "Repository size should be 1")
	assert.Equal(t, *entity, repoValidationsCodes.Items[0], "Should be same entity")
}

func TestShouldFindByUUIDValidationCodeInstance(t *testing.T) {
	setupTestValidationsCodes()

	code := "000000"
	now := time.Now().Format(time.RFC3339)
	entity, _ := entities.NewValidationCode(&code, &now, nil, nil)

	repoValidationsCodes.Create(*entity, nil)

	_, err := repoValidationsCodes.FindByUUID(uuid.New())

	assert.NotNil(t, err, "Should not found a entity by UUID")

	schedule, err := repoValidationsCodes.FindByUUID(entity.GetID())
	assert.Nil(t, err, "Should found a entity by UUID")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldFindByIDValidationCodeInstance(t *testing.T) {
	setupTestValidationsCodes()

	code := "000000"
	now := time.Now().Format(time.RFC3339)
	entity, _ := entities.NewValidationCode(&code, &now, nil, nil)

	repoValidationsCodes.Create(*entity, nil)

	_, err := repoValidationsCodes.FindByID("")

	assert.NotNil(t, err, "Should not found a entity by string")

	schedule, err := repoValidationsCodes.FindByID(entity.GetID().String())
	assert.Nil(t, err, "Should found a entity by string")

	assert.Equal(t, entity, schedule, "Entities are the same")
}

func TestShouldDeleteValidationCodeInstance(t *testing.T) {
	setupTestValidationsCodes()

	code := "000000"
	now := time.Now().Format(time.RFC3339)
	entity, _ := entities.NewValidationCode(&code, &now, nil, nil)

	repoValidationsCodes.Create(*entity, nil)

	err := repoValidationsCodes.Delete(*entity)

	assert.Nil(t, err, "Should not throw an error on delete entity")
	_, err = repoValidationsCodes.FindByUUID(entity.GetID())

	assert.NotNil(t, err, "Should throw an error on find entity")
}

func TestShouldDeleteManyValidationCodesInstance(t *testing.T) {
	setupTestValidationsCodes()

	var entitiesToCreate []entities.ValidationCode

	for i := 0; i < 10; i++ {
		code := strings.Repeat(strconv.Itoa(i), 6)
		now := time.Now().Format(time.RFC3339)
		entity, _ := entities.NewValidationCode(&code, &now, nil, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoValidationsCodes.CreateMany(entitiesToCreate)

	err := repoValidationsCodes.DeleteMany(entitiesToCreate)

	assert.Nil(t, err, "should not throw an error")

	_, err = repoValidationsCodes.FindByUUID(entitiesToCreate[0].GetID())
	assert.NotNil(t, err, "should throw an error on find by id")
}

func TestShouldHandleCurrentUserValidationCode(t *testing.T) {
	setupTestValidationsCodes()

	var entitiesToCreate []entities.ValidationCode

	user := entities.User{Entity: ent.NewEntity(nil, nil)}

	for i := 0; i < 10; i++ {
		code := strings.Repeat(strconv.Itoa(i), 6)
		now := time.Now().Format(time.RFC3339)
		entity, _ := entities.NewValidationCode(&code, &now, &user, nil)
		entitiesToCreate = append(entitiesToCreate, *entity)
	}

	repoValidationsCodes.CreateMany(entitiesToCreate)

	foundedEntity, err := repoValidationsCodes.FindCurrentUserValidationCode(user)

	assert.NotNil(t, err, "Should have an error after not found future validation code")
	assert.Nil(t, foundedEntity, "Should have no entity after not found future validation code")

	code := strings.Repeat("a", 6)
	now := time.Now().Add(4 * time.Minute).Format(time.RFC3339)
	entity, _ := entities.NewValidationCode(&code, &now, &user, nil)
	repoValidationsCodes.Create(*entity, nil)
	foundedEntity, err = repoValidationsCodes.FindCurrentUserValidationCode(user)

	assert.Nil(t, err, "Should have no error on entity with future expiration time")
	assert.NotNil(t, foundedEntity, "Should have a entity")

}
