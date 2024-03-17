package inmemory

import (
	errors2 "errors"
	"health/core/application/errors"
	"health/core/clients/domain/entities"

	"time"

	"github.com/google/uuid"
)

type ValidationsCodesInMemoryRepository struct {
	Items []entities.ValidationCode
}

func (repo *ValidationsCodesInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *ValidationsCodesInMemoryRepository) FindCurrentUserValidationCode(user entities.User) (*entities.ValidationCode, error) {
	now := time.Now()
	for _, item := range repo.Items {
		userRepo := item.GetPerson()
		if userRepo.GetID() == user.GetID() && item.DeletedAt == nil && item.GetExpirationDate().After(now) {
			return &item, nil
		}
	}

	return nil, errors.NewNotFoundError("Could not found current validation code for user " + user.GetID().String())
}

func (repo *ValidationsCodesInMemoryRepository) FindByUUID(id uuid.UUID) (*entities.ValidationCode, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *ValidationsCodesInMemoryRepository) FindByID(id string) (*entities.ValidationCode, error) {

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

func (repo *ValidationsCodesInMemoryRepository) Create(entity entities.ValidationCode, _ *any) error {
	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *ValidationsCodesInMemoryRepository) CreateMany(entity []entities.ValidationCode) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *ValidationsCodesInMemoryRepository) Update(entity entities.ValidationCode) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *ValidationsCodesInMemoryRepository) Delete(entity entities.ValidationCode) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *ValidationsCodesInMemoryRepository) DeleteMany(entitiesToDelete []entities.ValidationCode) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.ValidationCode{}

	for _, item := range entitiesToDelete {
		if index := repo.FindIndex(item.GetID()); index == -1 {
			indexesNotFound = append(indexesNotFound, item.GetID())
		} else {
			indexesFound = append(indexesFound, item)
		}
	}

	for _, item := range indexesFound {
		repo.Delete(item)
	}

	return nil
}

func NewValidationsCodesInMemoryRepository() *ValidationsCodesInMemoryRepository {
	return &ValidationsCodesInMemoryRepository{}
}
