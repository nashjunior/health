package repositories

import (
	"fmt"
	"health/core/application/errors"
	repo "health/core/application/repositories"
	valueobjects "health/core/application/value-objects"
	"math/big"

	ent "health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"health/core/clients/infra/db/gorm/entities"
	"health/core/infra/db/gorm"
	"time"

	"github.com/google/uuid"

	orm "gorm.io/gorm"
)

type TransactionsGorm struct {
	db           *orm.DB
	transaction  *entities.Transaction
	transactions *[]entities.Transaction
}

func (repo *TransactionsGorm) FindAllByTypeTransaction(id string, conn repo.IConnection) []ent.Transaction {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	repo.db.Joins("TypeTransaction", "TypeTransaction.uuid = ?", id).Find(&repo.transactions)
	time.Sleep(40 * time.Second)

	var entities []ent.Transaction

	for _, item := range *repo.transactions {
		uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.TypeTransaction.Uuid})
		accountType := ent.AccountType(item.TypeTransaction.AccountType)
		operationType := ent.OperationType(item.TypeTransaction.OperationType)

		typeT, err := ent.NewTypeTransaction(&item.TypeTransaction.Name, &accountType, &operationType, &uniqueTypeTransactionId)

		if err == nil {
			uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

			typeT.SetInternalId(item.TypeTransaction.Id)

			ent, err := ent.NewTransaction(
				&item.Date,
				&item.Value,
				typeT,
				&uniqueId,
			)

			if err == nil {
				entities = append(entities, *ent)
			}
		}
	}
	return entities
}

func (repo *TransactionsGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.Transaction, error) {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("Id").Joins("TypeTransaction").First(&repo.transaction, "transactions.uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.transaction == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	uniqueTypeTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.transaction.TypeTransaction.Uuid})
	accountType := ent.AccountType(repo.transaction.TypeTransaction.AccountType)
	operationType := ent.OperationType(repo.transaction.TypeTransaction.OperationType)

	typeT, err := ent.NewTypeTransaction(&repo.transaction.TypeTransaction.Name, &accountType, &operationType, &uniqueTypeTransactionId)

	if err != nil {
		return nil, fmt.Errorf("Could not parse type transaction for transaction")
	}

	uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.transaction.Uuid})

	return ent.NewTransaction(
		&repo.transaction.Date,
		&repo.transaction.Value,
		typeT,
		&uniqueId,
	)

}

func (repo *TransactionsGorm) FindByID(id string, conn repo.IConnection) (*ent.Transaction, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)

}

func (repo *TransactionsGorm) Find(params *repositories.SearchParamTransactions, conn repo.IConnection) []ent.Transaction {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("id").Find(&repo.transactions); result.Error != nil {
		return []ent.Transaction{}
	}

	var entities []ent.Transaction

	for _, item := range *repo.transactions {
		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		ent, err := ent.NewTransaction(
			&item.Date,
			&item.Value,
			nil,
			&uniqueId,
		)

		if err == nil {
			entities = append(entities, *ent)
		}

	}

	return entities
}

func (repo *TransactionsGorm) FindAndCount(params *repositories.SearchParamTransactions, conn repo.IConnection) repositories.IResponseSearchableTransactions {
	items := repo.Find(params, conn)

	return repositories.IResponseSearchableTransactions{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *TransactionsGorm) Create(entity *ent.Transaction, conn repo.IConnection) error {

	typeTransaction := entity.GetTypeTransaction()

	ent := entities.Transaction{
		Uuid:              entity.GetID(),
		Date:              entity.GetDate(),
		Value:             entity.GetValue(),
		IdTypeTransaction: typeTransaction.GetInternalId(),
	}

	ommitFields := []string{"CreatedAt", "DeletedAt", "UpdatedAt"}

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

func (repo *TransactionsGorm) CreateMany(items []ent.Transaction, conn repo.IConnection) error {

	var itemsEnt []entities.Transaction

	for _, item := range items {
		typeTransaction := item.GetTypeTransaction()

		ent := entities.Transaction{
			Uuid:              item.GetID(),
			Date:              item.GetDate(),
			Value:             item.GetValue(),
			IdTypeTransaction: typeTransaction.GetInternalId(),
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

func (repo *TransactionsGorm) Update(entity ent.Transaction, conn repo.IConnection) error {

	ent := entities.Transaction{
		Uuid:      entity.GetID(),
		Date:      entity.GetDate(),
		Value:     entity.GetValue(),
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	ommitedFields := []string{"CreatedAt", "DeletedAt"}

	if result := repo.db.Omit(ommitedFields...).Save(&ent); result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TransactionsGorm) Delete(entity ent.Transaction, conn repo.IConnection) error {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Debug().
		Where("uuid = ?", entity.GetID()).
		Delete(&entities.Transaction{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *TransactionsGorm) DeleteMany(entitiesToDelete []ent.Transaction, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.Transaction{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewTransactionsGorm() repositories.ITransactionsRepository {
	return &TransactionsGorm{}
}
