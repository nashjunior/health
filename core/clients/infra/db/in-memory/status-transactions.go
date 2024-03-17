package inmemory

import (
	errors2 "errors"
	"fmt"
	"health/core/application/errors"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"

	coreRepository "health/core/application/repositories"

	"math/big"
	"time"

	"github.com/google/uuid"
)

type StatusTransactionsInMemoryRepository struct {
	Items []entities.StatusTransaction
}

func (repo *StatusTransactionsInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *StatusTransactionsInMemoryRepository) FindAllByTransaction(id string, tx coreRepository.IConnection) []entities.StatusTransaction {
	var transactionsSlice []entities.StatusTransaction
	for _, item := range repo.Items {
		if transaction, isDeleted := item.GetTransaction(), item.DeletedAt == nil; transaction.GetID().String() == id && isDeleted {
			transactionsSlice = append(transactionsSlice, item)
		}
	}

	return transactionsSlice
}

func (repo *StatusTransactionsInMemoryRepository) FindACurrentStatusTransaction(id string, tx coreRepository.IConnection) (*entities.StatusTransaction, error) {

	for _, item := range repo.Items {
		if transaction, isDeleted := item.GetTransaction(), item.DeletedAt == nil; transaction.GetID().String() == id && isDeleted {
			return &item, nil
		}
	}

	return nil, errors.NewNotFoundError("Cound not found entity using id")
}

func (repo *StatusTransactionsInMemoryRepository) FindByUUID(id uuid.UUID, _ coreRepository.IConnection) (*entities.StatusTransaction, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *StatusTransactionsInMemoryRepository) FindByID(id string, _ coreRepository.IConnection) (*entities.StatusTransaction, error) {

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

func (repo *StatusTransactionsInMemoryRepository) filter(params *repositories.SearchParamStatusTransactions) []entities.StatusTransaction {

	// filteredItems := []entities.StatusTransaction{}

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

func (repo *StatusTransactionsInMemoryRepository) paginate(items []entities.StatusTransaction, params *repositories.SearchParamStatusTransactions) []entities.StatusTransaction {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.StatusTransaction{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *StatusTransactionsInMemoryRepository) Find(params *repositories.SearchParamStatusTransactions, _ coreRepository.IConnection) []entities.StatusTransaction {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *StatusTransactionsInMemoryRepository) FindAndCount(_ *repositories.SearchParamStatusTransactions, _ coreRepository.IConnection) repositories.IResponseSearchableStatusTransactions {
	return repositories.IResponseSearchableStatusTransactions{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *StatusTransactionsInMemoryRepository) Create(entity entities.StatusTransaction, tx coreRepository.IConnection) error {

	if tx != nil {
		fmt.Println("Create person with transaction")
	}

	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *StatusTransactionsInMemoryRepository) CreateMany(entity []entities.StatusTransaction, _ coreRepository.IConnection) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *StatusTransactionsInMemoryRepository) Update(entity entities.StatusTransaction, _ coreRepository.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *StatusTransactionsInMemoryRepository) Delete(entity entities.StatusTransaction, _ coreRepository.IConnection) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *StatusTransactionsInMemoryRepository) DeleteMany(entitiesToDelete []entities.StatusTransaction, tx coreRepository.IConnection) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.StatusTransaction{}

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

func NewStatusTransactionsInMemoryRepository() *StatusTransactionsInMemoryRepository {
	return &StatusTransactionsInMemoryRepository{}
}
