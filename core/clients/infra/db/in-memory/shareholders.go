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

type ShareholdersInMemoryRepository struct {
	Items []entities.Shareholder
}

func (repo *ShareholdersInMemoryRepository) FindIndex(id uuid.UUID) int {
	for index, item := range repo.Items {
		if item.GetID() == id && item.DeletedAt == nil {
			return index
		}
	}

	return -1
}

func (repo *ShareholdersInMemoryRepository) FindByUUID(id uuid.UUID, _ *any) (*entities.Shareholder, error) {
	index := repo.FindIndex(id)

	if index == -1 {
		return nil, errors.NewNotFoundError("Entity not found using UUID: " + id.String())
	}

	return &repo.Items[index], nil
}

func (repo *ShareholdersInMemoryRepository) FindByID(id string, _ *any) (*entities.Shareholder, error) {

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

func (repo *ShareholdersInMemoryRepository) filter(params *repositories.SearchParamShareholders) []entities.Shareholder {

	// filteredItems := []entities.Shareholder{}

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

func (repo *ShareholdersInMemoryRepository) paginate(items []entities.Shareholder, params *repositories.SearchParamShareholders) []entities.Shareholder {
	hasPagination := params.Pagination != nil

	if hasPagination {
		pagination := *params.Pagination
		startIndex := (pagination.Page - 1) * pagination.PerPage
		endIndex := pagination.Page * pagination.PerPage

		if startIndex >= len(items) {
			return []entities.Shareholder{}
		}

		if endIndex > len(items) {
			endIndex = len(items)
		}

		return items[startIndex:endIndex]
	}

	return items
}

func (repo *ShareholdersInMemoryRepository) Find(params *repositories.SearchParamShareholders, _ *any) []entities.Shareholder {
	if params == nil {
		return repo.Items
	}

	filteredItems := repo.filter(params)

	filteredItems = repo.paginate(filteredItems, params)

	return filteredItems
}

func (repo *ShareholdersInMemoryRepository) FindAndCount(params *repositories.SearchParamShareholders, _ *any) repositories.ResponseSearchableShareholders {
	return repositories.ResponseSearchableShareholders{
		Total: *big.NewInt(int64(len(repo.Items))),
		Items: repo.Items,
	}
}

func (repo *ShareholdersInMemoryRepository) FindAllByPerson(person entities.Person, _ *any) []entities.Shareholder {
	customersRoles := []entities.Shareholder{}

	for _, item := range repo.Items {
		if deleted, personCustomerRole := item.DeletedAt, item.GetOwner(); deleted == nil && personCustomerRole.GetID() == person.GetID() {
			customersRoles = append(customersRoles, item)
		}
	}
	return customersRoles
}

func (repo *ShareholdersInMemoryRepository) FindAllByEnterprises(enterprise entities.Enterprise, _ *any) []entities.Shareholder {
	customersRoles := []entities.Shareholder{}

	for _, item := range repo.Items {
		if deleted, shareEnterprise := item.DeletedAt, item.GetEnterprise(); deleted == nil && shareEnterprise.GetID() == enterprise.GetID() {
			customersRoles = append(customersRoles, item)
		}
	}
	return customersRoles
}

func (repo *ShareholdersInMemoryRepository) Create(entity entities.Shareholder, _ *any) error {
	repo.Items = append(repo.Items, entity)
	return nil
}

func (repo *ShareholdersInMemoryRepository) CreateMany(entity []entities.Shareholder, _ *any) error {
	repo.Items = append(repo.Items, entity...)
	return nil
}

func (repo *ShareholdersInMemoryRepository) Update(entity entities.Shareholder, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	repo.Items[index] = entity

	return nil
}

func (repo *ShareholdersInMemoryRepository) Delete(entity entities.Shareholder, _ *any) error {

	index := repo.FindIndex(entity.GetID())

	if index == -1 {
		return errors.NewNotFoundError("Entity not found using UUID: " + entity.GetID().String())
	}

	deletedAt := time.Now()
	entity.DeletedAt = &deletedAt

	repo.Items[index] = entity

	return nil
}

func (repo *ShareholdersInMemoryRepository) DeleteMany(entitiesToDelete []entities.Shareholder, tx *any) error {

	indexesNotFound := uuid.UUIDs{}
	indexesFound := []entities.Shareholder{}

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

func NewShareholdersInMemoryRepository() *ShareholdersInMemoryRepository {
	return &ShareholdersInMemoryRepository{}
}
