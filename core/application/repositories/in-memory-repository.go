package repositories

import (
	"health/core/application/entities"
	"health/core/application/errors"
	"time"

	"github.com/google/uuid"
)

type AbstractIntMemoryRepository struct {
	Items []entities.Entity
}

func (repo *AbstractIntMemoryRepository) findIndex(id uuid.UUID) *int {
	var index *int
	for i, item := range repo.Items {
		if id == item.UniqueEntityUUID.Id && item.DeletedAt != nil {
			index = &i
		}
	}
	return index
}

func (repo *AbstractIntMemoryRepository) FindByUUID(id uuid.UUID) (*entities.Entity, error) {
	index := repo.findIndex(id)

	if index == nil {
		return nil, errors.NewInvalidUUidError(id)
	}
	return &repo.Items[*index], nil
}

func (repo *AbstractIntMemoryRepository) FindByID(id string) (*entities.Entity, error) {

	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, &errors.InvalidUUidError{}
	}

	index := repo.findIndex(parsedId)

	if index == nil {
		return nil, errors.NewInvalidUUidError(parsedId)
	}
	return &repo.Items[*index], nil
}

func (repo *AbstractIntMemoryRepository) Create(entity entities.Entity) error {
	repo.Items = append(repo.Items, entity)

	return nil
}

func (repo *AbstractIntMemoryRepository) CreateMany(entities []entities.Entity) error {
	repo.Items = append(repo.Items, entities...)

	return nil
}

func (repo *AbstractIntMemoryRepository) update(entity entities.Entity) error {

	index := repo.findIndex(entity.UniqueEntityUUID.Id)

	if index == nil {
		return errors.NewInvalidUUidError(entity.UniqueEntityUUID.Id)
	}

	entity.UpdatedAt = &time.Time{}

	repo.Items[*index] = entity

	return nil

}

func (repo *AbstractIntMemoryRepository) Delete(entity entities.Entity) error {

	index := repo.findIndex(entity.UniqueEntityUUID.Id)

	if index == nil {
		return errors.NewInvalidUUidError(entity.UniqueEntityUUID.Id)
	}

	entity.DeletedAt = &time.Time{}

	repo.Items[*index] = entity

	return nil

}

func (repo *AbstractIntMemoryRepository) DeleteMany(entity entities.Entity) error {

	index := repo.findIndex(entity.UniqueEntityUUID.Id)

	if index == nil {
		return errors.NewInvalidUUidError(entity.UniqueEntityUUID.Id)
	}

	entity.DeletedAt = &time.Time{}

	repo.Items[*index] = entity

	return nil

}
