package inmemory

import (
	errors2 "errors"
	"fmt"
	"health/core/application/errors"
	"health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"

	"math/big"
	"time"

	"github.com/google/uuid"
)

type PersonsInMemoryRepository struct {
	Items []entities.Person
}

func (repo *PersonsInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *PersonsInMemoryRepository) FindByUUID(id uuid.UUID, _ *any) (*entities.Person, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *PersonsInMemoryRepository) FindByID(id string, _ *any) (*entities.Person, error) {

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

func (repo *PersonsInMemoryRepository) filter(params *repositories.SearchParamPersons) []entities.Person {

	// filteredItems := []entities.Person{}

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

func (repo *PersonsInMemoryRepository) paginate(items []entities.Person, params *repositories.SearchParamPersons) []entities.Person {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.Person{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *PersonsInMemoryRepository) Find(params *repositories.SearchParamPersons, _ *any) []entities.Person {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *PersonsInMemoryRepository) FindAndCount(_ *repositories.SearchParamPersons, _ *any) repositories.ResponseSearchablePersons {
	return repositories.ResponseSearchablePersons{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *PersonsInMemoryRepository) FindByUser(user entities.User, _ *any) (*entities.Person, error) {

	for _, item := range repo.Items {
		if deletado, userPerson := item.DeletedAt, item.GetUser(); deletado == nil && userPerson.GetID() == user.GetID() {
			return &item, nil
		}
	}

	return nil, &errors.InvalidUUidError{Id: user.GetID()}
}

func (repo *PersonsInMemoryRepository) Create(entity entities.Person, tx *any) error {

	if tx != nil {
		fmt.Println("Create person with transaction")
	}

	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *PersonsInMemoryRepository) CreateMany(entity []entities.Person, _ *any) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *PersonsInMemoryRepository) Update(entity entities.Person, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *PersonsInMemoryRepository) Delete(entity entities.Person, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *PersonsInMemoryRepository) DeleteMany(entitiesToDelete []entities.Person, tx *any) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.Person{}

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

func NewPersonsInMemoryRepository() *PersonsInMemoryRepository {
	return &PersonsInMemoryRepository{}
}
