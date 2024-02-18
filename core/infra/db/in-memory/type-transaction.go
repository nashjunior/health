package inmemory

import (
	errors2 "errors"
	"fmt"
	"health/core/application/errors"
	repo "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"

	"math/big"
	"time"

	"github.com/google/uuid"
)

type TypeTransactionsInMemoryRepository struct {
	Items []entities.TypeTransaction
}

func (repo *TypeTransactionsInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *TypeTransactionsInMemoryRepository) FindByUUID(id uuid.UUID, _ *repo.IConnection) (*entities.TypeTransaction, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *TypeTransactionsInMemoryRepository) FindByID(id string, _ *repo.IConnection) (*entities.TypeTransaction, error) {

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

func (repo *TypeTransactionsInMemoryRepository) filter(params *repositories.SearchParamPersons) []entities.TypeTransaction {

	// filteredItems := []entities.TypeTransaction{}

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

func (repo *TypeTransactionsInMemoryRepository) paginate(items []entities.TypeTransaction, params *repositories.SearchParamPersons) []entities.TypeTransaction {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.TypeTransaction{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *TypeTransactionsInMemoryRepository) Find(params *repositories.SearchParamPersons, _ *repo.IConnection) []entities.TypeTransaction {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *TypeTransactionsInMemoryRepository) FindAndCount(_ *repositories.SearchParamTypeTransactions, _ *repo.IConnection) repositories.IResponseSearchableTypeTransactions {
	return repositories.IResponseSearchableTypeTransactions{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *TypeTransactionsInMemoryRepository) CreateEntity(entity entities.TypeTransaction, tx *repo.IConnection) error {

	if tx != nil {
		fmt.Println("Create person with transaction")
	}

	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *TypeTransactionsInMemoryRepository) CreateMany(entity []entities.TypeTransaction, _ *repo.IConnection) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *TypeTransactionsInMemoryRepository) Update(entity entities.TypeTransaction, _ *repo.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *TypeTransactionsInMemoryRepository) Delete(entity entities.TypeTransaction, _ *repo.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *TypeTransactionsInMemoryRepository) DeleteMany(entitiesToDelete []entities.TypeTransaction, tx *repo.IConnection) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.TypeTransaction{}

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

func NewTypeTransactionsInMemoryRepository() *TypeTransactionsInMemoryRepository {
	return &TypeTransactionsInMemoryRepository{}
}
