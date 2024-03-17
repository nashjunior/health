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

type TransactionsInMemoryRepository struct {
	Items []entities.Transaction
}

func (repo *TransactionsInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *TransactionsInMemoryRepository) FindAllByTypeTransaction(id string, tx repo.IConnection) []entities.Transaction {
	var transactionsSlice []entities.Transaction
	for _, item := range repo.Items {
		if typeTransaction, isDeleted := item.GetTypeTransaction(), item.DeletedAt == nil; typeTransaction.GetID().String() == id && isDeleted {
			transactionsSlice = append(transactionsSlice, item)
		}
	}

	return transactionsSlice
}

func (repo *TransactionsInMemoryRepository) FindByUUID(id uuid.UUID, _ repo.IConnection) (*entities.Transaction, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *TransactionsInMemoryRepository) FindByID(id string, _ repo.IConnection) (*entities.Transaction, error) {

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

func (repo *TransactionsInMemoryRepository) filter(params *repositories.SearchParamPersons) []entities.Transaction {

	// filteredItems := []entities.Transaction{}

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

func (repo *TransactionsInMemoryRepository) paginate(items []entities.Transaction, params *repositories.SearchParamPersons) []entities.Transaction {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.Transaction{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *TransactionsInMemoryRepository) Find(params *repositories.SearchParamPersons, _ repo.IConnection) []entities.Transaction {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *TransactionsInMemoryRepository) FindAndCount(_ *repositories.SearchParamTransactions, _ repo.IConnection) repositories.IResponseSearchableTransactions {
	return repositories.IResponseSearchableTransactions{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *TransactionsInMemoryRepository) Create(entity entities.Transaction, tx repo.IConnection) error {

	if tx != nil {
		fmt.Println("Create person with transaction")
	}

	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *TransactionsInMemoryRepository) CreateMany(entity []entities.Transaction, _ repo.IConnection) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *TransactionsInMemoryRepository) Update(entity entities.Transaction, _ repo.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *TransactionsInMemoryRepository) Delete(entity entities.Transaction, _ repo.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *TransactionsInMemoryRepository) DeleteMany(entitiesToDelete []entities.Transaction, tx repo.IConnection) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.Transaction{}

	for _, item := range entitiesToDelete {
		if index := repo.FindIndex(item.GetID()); index == -1 {
			indexesNotFound = append(indexesNotFound, item.GetID())
		} else {
			indexesFound = append(indexesFound, item)
		}
	}

	for _, item := range indexesFound {
		repo.Delete(item, nil)
	}

	return nil
}

func NewTransactionsInMemoryRepository() *TransactionsInMemoryRepository {
	return &TransactionsInMemoryRepository{}
}
