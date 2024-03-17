package inmemory

import (
	errors2 "errors"
	"fmt"
	"health/core/application/errors"
	coreRepository "health/core/application/repositories"
	"health/health/domain/entities"
	"health/health/domain/repositories"
	"strings"

	"math/big"
	"time"

	"github.com/google/uuid"
)

type DiseasesInMemoryRepository struct {
	Items []entities.Disease
}

func (repo *DiseasesInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *DiseasesInMemoryRepository) FindByUUID(id uuid.UUID, _ coreRepository.IConnection) (*entities.Disease, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *DiseasesInMemoryRepository) FindByID(id string, _ coreRepository.IConnection) (*entities.Disease, error) {

	uuid, err := uuid.Parse(id)

	if err != nil {
		return nil, errors2.New("Could not parse id " + id)
	}

	index := repo.FindIndex(uuid)

	if index == -1 {
		fmt.Println("aqui")
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id)
	}

	return &repo.Items[index], nil
}

func (repo *DiseasesInMemoryRepository) FindByName(name string, tx coreRepository.IConnection) (*entities.Disease, error) {
	index := -1

	for idx, item := range repo.Items {
		if strings.Contains(strings.ToLower(item.GetName()), strings.ToLower(name)) && item.DeletedAt == nil {
			index = idx
		}
	}

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using name: " + name)
	}

	return &repo.Items[index], nil
}

func (repo *DiseasesInMemoryRepository) filter(params *repositories.SearchParamDisease) []entities.Disease {

	// filteredItems := []entities.Disease{}

	// for _, item := range repo.Items {

	// 	if query := params.Query; query != nil {
	// 		fields := query.Fields
	// 		value := query.Value

	// 		containsValue := false
	// 		for _, field := range fields {
	// 			switch strings.ToLower(field) {
	// 			case "name":
	// 				containsValue = strings.Contains(item.GetID().String(), value)
	// 			default:

	// 			}
	// 		}

	// 	}

	// }

	// hasFilter := params.Query != nil
	// if hasFilter {
	// 	return filteredItems
	// }

	return repo.Items
}

func (repo *DiseasesInMemoryRepository) paginate(items []entities.Disease, params *repositories.SearchParamDisease) []entities.Disease {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.Disease{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *DiseasesInMemoryRepository) Find(params *repositories.SearchParamDisease, _ coreRepository.IConnection) []entities.Disease {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *DiseasesInMemoryRepository) FindAndCount(params *repositories.SearchParamDisease, _ coreRepository.IConnection) repositories.IResponseSearchableDiseases {
	return repositories.IResponseSearchableDiseases{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *DiseasesInMemoryRepository) Create(entity *entities.Disease, _ coreRepository.IConnection) error {
	repo.Items = append(repo.Items, *entity)
	return nil
}

func (repo *DiseasesInMemoryRepository) CreateMany(entity []entities.Disease, _ coreRepository.IConnection) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *DiseasesInMemoryRepository) Update(entity entities.Disease, _ coreRepository.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *DiseasesInMemoryRepository) Delete(entity entities.Disease, _ coreRepository.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *DiseasesInMemoryRepository) DeleteMany(entitiesToDelete []entities.Disease, tx coreRepository.IConnection) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.Disease{}

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

func NewDiseasesInMemoryRepository() repositories.IDiseasesRepository {
	return &DiseasesInMemoryRepository{}
}
