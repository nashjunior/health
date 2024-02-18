package repositories

import (
	"health/core/application/errors"
	valueobjects "health/core/application/value-objects"
	"health/core/infra/db/gorm"
	"health/core/infra/db/gorm/entities"
	"math/big"

	repo "health/core/application/repositories"
	ent "health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"

	"github.com/google/uuid"
	orm "gorm.io/gorm"
)

type TypesTransactionsGorm struct {
	db *orm.DB

	typeTransaction   *entities.TypeTransaction
	typesTransactions *[]entities.TypeTransaction
}

func (repo *TypesTransactionsGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.TypeTransaction, error) {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.First(&repo.typeTransaction, "uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.typeTransaction == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.typeTransaction.Uuid})

	accountType := ent.AccountType(repo.typeTransaction.AccountType)
	operationType := ent.OperationType(repo.typeTransaction.OperationType)

	entity, err := ent.NewTypeTransaction(
		&repo.typeTransaction.Name,
		&accountType,
		&operationType,
		&uniqueId,
	)

	if err != nil {
		return nil, err
	}

	entity.SetInternalId(repo.typeTransaction.Id)

	return entity, nil

}

func (repo *TypesTransactionsGorm) FindByID(id string, conn repo.IConnection) (*ent.TypeTransaction, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)
}

func (repo *TypesTransactionsGorm) Find(params *repositories.SearchParamPersons, conn repo.IConnection) []ent.TypeTransaction {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Find(&repo.typesTransactions); result.Error != nil {
		return []ent.TypeTransaction{}
	}

	var entities []ent.TypeTransaction

	for _, item := range *repo.typesTransactions {
		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		accountType := ent.AccountType(item.AccountType)
		operationType := ent.OperationType(item.OperationType)

		ent, err := ent.NewTypeTransaction(
			&item.Name,
			&accountType,
			&operationType,
			&uniqueId,
		)

		if err == nil {
			ent.SetInternalId(repo.typeTransaction.Id)
			entities = append(entities, *ent)
		}
	}

	return entities

}

func (repo *TypesTransactionsGorm) FindAndCount(_ *repositories.SearchParamTypeTransactions, conn repo.IConnection) repositories.IResponseSearchableTypeTransactions {

	items := repo.Find(nil, conn)
	return repositories.IResponseSearchableTypeTransactions{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *TypesTransactionsGorm) Create(entity *ent.TypeTransaction, conn repo.IConnection) error {

	ent := entities.TypeTransaction{
		Uuid:          entity.GetID(),
		Name:          entity.GetName(),
		AccountType:   uint8(entity.GetAccountType()),
		OperationType: uint8(entity.GetOperationType()),
	}

	ommitFields := []string{"CreatedAt", "Id", "DeletedAt", "UpdatedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit(ommitFields...).Create(&ent); result.Error != nil {
		return result.Error
	}

	entity.SetInternalId(ent.Id)

	return nil
}

func (repo *TypesTransactionsGorm) CreateMany(items []ent.TypeTransaction, conn repo.IConnection) error {

	var itemsEnt []entities.TypeTransaction

	for _, item := range items {
		ent := entities.TypeTransaction{
			Uuid:          item.GetID(),
			Name:          item.GetName(),
			AccountType:   uint8(item.GetAccountType()),
			OperationType: uint8(item.GetOperationType()),
		}

		itemsEnt = append(itemsEnt, ent)
	}

	ommitFields := []string{"CreatedAt", "Id", "DeletedAt", "UpdatedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit(ommitFields...).Create(&itemsEnt); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TypesTransactionsGorm) Update(entity ent.TypeTransaction, conn repo.IConnection) error {

	ent := entities.TypeTransaction{
		Uuid:          entity.GetID(),
		Name:          entity.GetName(),
		AccountType:   uint8(entity.GetAccountType()),
		OperationType: uint8(entity.GetOperationType()),
		CreatedAt:     entity.CreatedAt,
		UpdatedAt:     entity.UpdatedAt,
	}

	ommitedFields := []string{"CreatedAt", "DeletedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit(ommitedFields...).Save(&ent); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TypesTransactionsGorm) Delete(entity ent.TypeTransaction, conn repo.IConnection) error {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Delete(&entities.TypeTransaction{}, entity.GetID()); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TypesTransactionsGorm) DeleteMany(entitiesToDelete []ent.TypeTransaction, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.TypeTransaction{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewTypesTransactionsGorm() repositories.ITypeTransactionsRepository {
	return &TypesTransactionsGorm{}
}
