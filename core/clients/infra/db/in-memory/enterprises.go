package inmemory

import (
	errors2 "errors"
	"health/core/application/errors"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"

	"math/big"
	"time"

	"github.com/google/uuid"
)

type EnterprisesInMemoryRepository struct {
	Items []entities.Enterprise
}

func (repo *EnterprisesInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *EnterprisesInMemoryRepository) FindByUUID(id uuid.UUID, _ *any) (*entities.Enterprise, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *EnterprisesInMemoryRepository) FindByID(id string, _ *any) (*entities.Enterprise, error) {

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

func (repo *EnterprisesInMemoryRepository) filter(params *repositories.SearchParamEnterprises) []entities.Enterprise {

	// filteredItems := []entities.Enterprise{}

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

func (repo *EnterprisesInMemoryRepository) paginate(items []entities.Enterprise, params *repositories.SearchParamEnterprises) []entities.Enterprise {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.Enterprise{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *EnterprisesInMemoryRepository) Find(params *repositories.SearchParamEnterprises, _ *any) []entities.Enterprise {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *EnterprisesInMemoryRepository) FindAndCount(params *repositories.SearchParamEnterprises, _ *any) repositories.ResponseSearchableEnterprises {
	return repositories.ResponseSearchableEnterprises{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *EnterprisesInMemoryRepository) Create(entity entities.Enterprise, _ *any) error {
	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *EnterprisesInMemoryRepository) CreateMany(entity []entities.Enterprise, _ *any) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *EnterprisesInMemoryRepository) Update(entity entities.Enterprise, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *EnterprisesInMemoryRepository) Delete(entity entities.Enterprise, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *EnterprisesInMemoryRepository) DeleteMany(entitiesToDelete []entities.Enterprise, tx *any) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.Enterprise{}

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

func NewEnterprisesInMemoryRepository() *EnterprisesInMemoryRepository {
	return &EnterprisesInMemoryRepository{}
}
