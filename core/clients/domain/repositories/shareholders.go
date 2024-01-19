package repositories

import (
	coreRepository "health/core/application/repositories"
	"health/core/clients/domain/entities"
	"math/big"

	"github.com/google/uuid"
)

type ResponseSearchableShareholders struct {
	Total big.Int
	Items []entities.Shareholder
}

type SearchParamShareholders struct {
	Query      *coreRepository.SearchableQuery
	Pagination *coreRepository.SearchablePagination
}

type IShareholdersRepository interface {
	FindByUUID(id uuid.UUID, tx *any) (*entities.Shareholder, error)
	FindByID(id string, tx *any) (*entities.Shareholder, error)
	//Find(params *SearchParamPharmacies) []entities.Shareholder
	//FindAndCount(params *SearchParamPharmacies) IResponseSearchablePharmacies

	FindAllByPerson(person entities.Person, tx *any) []entities.Shareholder
	FindAllByEnterprises(person entities.Enterprise, tx *any) []entities.Shareholder

	Create(entitiy entities.Shareholder, tx *any) error
	//CreateMany(entitiy []entities.Shareholder) error

	Update(entity entities.Shareholder, tx *any) error

	Delete(entity entities.Shareholder, tx *any) error
	DeleteMany(entidades []entities.Shareholder, tx *any) error
}
