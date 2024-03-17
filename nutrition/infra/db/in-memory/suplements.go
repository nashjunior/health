package inmemory

import (
	errors2 "errors"
	"fmt"
	"health/core/application/errors"
	coreRepository "health/core/application/repositories"
	"health/nutrition/domain/entities"
	"health/nutrition/domain/repositories"
	"strings"

	"math/big"
	"time"

	"github.com/google/uuid"
)

type SuplementsInMemoryRepository struct {
	Items []entities.Suplement
}

func (repo *SuplementsInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *SuplementsInMemoryRepository) FindByUUID(id uuid.UUID, _ coreRepository.IConnection) (*entities.Suplement, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *SuplementsInMemoryRepository) FindByID(id string, _ coreRepository.IConnection) (*entities.Suplement, error) {

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

func (repo *SuplementsInMemoryRepository) FindByName(name string, tx coreRepository.IConnection) (*entities.Suplement, error) {
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

func (repo *SuplementsInMemoryRepository) filter(_ *repositories.SearchParamEquipament) []entities.Suplement {

	// filteredItems := []entities.Suplement{}

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

func (repo *SuplementsInMemoryRepository) paginate(items []entities.Suplement, params *repositories.SearchParamEquipament) []entities.Suplement {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.Suplement{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *SuplementsInMemoryRepository) Find(params *repositories.SearchParamEquipament, _ coreRepository.IConnection) []entities.Suplement {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *SuplementsInMemoryRepository) FindAndCount(params *repositories.SearchParamEquipament, _ coreRepository.IConnection) repositories.IResponseSearchableSuplement {
	return repositories.IResponseSearchableSuplement{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *SuplementsInMemoryRepository) Create(entity *entities.Suplement, _ coreRepository.IConnection) error {
	repo.Items = append(repo.Items, *entity)
	return nil
}

func (repo *SuplementsInMemoryRepository) CreateMany(entity []entities.Suplement, _ coreRepository.IConnection) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *SuplementsInMemoryRepository) Update(entity entities.Suplement, _ coreRepository.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *SuplementsInMemoryRepository) Delete(entity entities.Suplement, _ coreRepository.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *SuplementsInMemoryRepository) DeleteMany(entitiesToDelete []entities.Suplement, tx coreRepository.IConnection) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.Suplement{}

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

func NewSuplementsInMemoryRepository() repositories.ISuplementsRepository {
	return &SuplementsInMemoryRepository{}
}
