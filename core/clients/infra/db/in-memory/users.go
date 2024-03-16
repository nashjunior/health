package inmemory

import (
	errors2 "errors"
	"health/core/application/errors"
	"health/core/clients/domain/entities"

	"time"

	"github.com/google/uuid"
)

type UsersInMemoryRepository struct {
	Items []entities.User
}

func (repo *UsersInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *UsersInMemoryRepository) FindByUUID(id uuid.UUID, _ *any) (*entities.User, error) {

	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *UsersInMemoryRepository) FindByID(id string, _ *any) (*entities.User, error) {

	uuid, err := uuid.Parse(id)

	if err != nil {
		return nil, errors2.New("Could not parse id " + id)
	}

	index := repo.FindIndex(uuid)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id)
	}

	return &repo.Items[index], nil
}

func (repo *UsersInMemoryRepository) Create(entity entities.User, _ *any) error {
	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *UsersInMemoryRepository) CreateMany(entity []entities.User, _ *any) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *UsersInMemoryRepository) Update(entity entities.User, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *UsersInMemoryRepository) Delete(entity entities.User, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *UsersInMemoryRepository) DeleteMany(entitiesToDelete []entities.User, tx *any) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.User{}

	for _, item := range entitiesToDelete {
		if index := repo.FindIndex(item.GetID()); index == -1 {
			indexesNotFound = append(indexesNotFound, item.GetID())
		} else {
			indexesFound = append(indexesFound, item)
		}
	}

	for _, item := range indexesFound {
		repo.Delete(item, tx)
	}

	return nil
}

func NewUsersInMemoryRepository() *UsersInMemoryRepository {
	return &UsersInMemoryRepository{}
}
