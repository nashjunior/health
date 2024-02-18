package repositories

import (
	"fmt"
	"health/core/application/errors"
	repo "health/core/application/repositories"
	valueobjects "health/core/application/value-objects"
	"math/big"

	ent "health/core/clients/domain/entities"
	"health/core/clients/domain/repositories"
	"health/core/infra/db/gorm"
	"health/core/infra/db/gorm/entities"

	"github.com/google/uuid"
	orm "gorm.io/gorm"
)

type StatusTransactionsGorm struct {
	db *orm.DB

	statusTransaction  *entities.StatusTransaction
	statusTransactions *[]entities.StatusTransaction
}

func (repo *StatusTransactionsGorm) FindAllByTransaction(id string, conn repo.IConnection) []ent.StatusTransaction {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Joins("Transaction", "Transaction.uuid = ?", id).Find(&repo.statusTransactions); result.Error != nil {
		return []ent.StatusTransaction{}
	}

	var entities []ent.StatusTransaction

	for _, item := range *repo.statusTransactions {
		uniqueTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Transaction.Uuid})

		typeT, err := ent.NewTransaction(&item.Transaction.Date, &item.Transaction.Value, nil, &uniqueTransactionId)

		if err == nil {
			uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

			fmt.Println(item.Transaction.Id)

			typeT.SetInternalId(item.Transaction.Id)

			status := ent.Status(item.Status)

			ent, err := ent.NewStatusTransaction(
				&status,
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

func (repo *StatusTransactionsGorm) FindACurrentStatusTransaction(id string, conn repo.IConnection) (*ent.StatusTransaction, error) {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		fmt.Println(gorm.Connection)
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Debug().Joins("Transaction", "Transaction.uuid = ?", id).Find(&repo.statusTransaction); result.Error != nil {
		return nil, result.Error
	}

	if repo.statusTransaction == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	uniqueTransactionId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.statusTransaction.Transaction.Uuid})

	typeT, err := ent.NewTransaction(&repo.statusTransaction.Transaction.Date, &repo.statusTransaction.Transaction.Value, nil, &uniqueTransactionId)

	if err != nil {
		return nil, fmt.Errorf("Could not parse transaction on find current transaction status")
	}

	uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.statusTransaction.Uuid})

	status := ent.Status(repo.statusTransaction.Status)

	return ent.NewStatusTransaction(
		&status,
		typeT,
		&uniqueId,
	)
}

func (repo *StatusTransactionsGorm) FindByUUID(id uuid.UUID, conn repo.IConnection) (*ent.StatusTransaction, error) {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("Id").First(&repo.statusTransaction, "uuid = ?", id); result.Error != nil {
		return nil, result.Error
	}

	if repo.statusTransaction == nil {
		return nil, errors.NewNotFoundError("Could not found error using id")
	}

	uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: repo.statusTransaction.Uuid})

	status := ent.Status(repo.statusTransaction.Status)

	return ent.NewStatusTransaction(
		&status,
		nil,
		&uniqueId,
	)

}

func (repo *StatusTransactionsGorm) FindByID(id string, conn repo.IConnection) (*ent.StatusTransaction, error) {
	parsedId, err := uuid.Parse(id)

	if err != nil {
		return nil, err
	}

	return repo.FindByUUID(parsedId, conn)

}

func (repo *StatusTransactionsGorm) Find(params *repositories.SearchParamStatusTransactions, conn repo.IConnection) []ent.StatusTransaction {

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Omit("Id").Find(&repo.statusTransactions); result.Error != nil {
		return []ent.StatusTransaction{}
	}

	var entities []ent.StatusTransaction

	for _, item := range *repo.statusTransactions {
		uniqueId := valueobjects.NewUniqueUUID(uuid.NullUUID{UUID: item.Uuid})

		status := ent.Status(item.Status)

		ent, err := ent.NewStatusTransaction(
			&status,
			nil,
			&uniqueId,
		)

		if err == nil {
			entities = append(entities, *ent)
		}

	}

	return entities
}

func (repo *StatusTransactionsGorm) FindAndCount(params *repositories.SearchParamStatusTransactions, conn repo.IConnection) repositories.IResponseSearchableStatusTransactions {
	items := repo.Find(params, conn)

	return repositories.IResponseSearchableStatusTransactions{
		Total: *big.NewInt(int64(len(items))),
		Items: items,
	}
}

func (repo *StatusTransactionsGorm) Create(entity *ent.StatusTransaction, conn repo.IConnection) error {

	transaction := entity.GetTransaction()

	ent := entities.StatusTransaction{
		Uuid:          entity.GetID(),
		Status:        uint8(entity.GetStatus()),
		IdTransaction: transaction.GetInternalId(),
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

func (repo *StatusTransactionsGorm) CreateMany(items []ent.StatusTransaction, conn repo.IConnection) error {

	var itemsEnt []entities.StatusTransaction

	for _, item := range items {
		transaction := item.GetTransaction()

		ent := entities.StatusTransaction{
			Uuid:          item.GetID(),
			Status:        uint8(item.GetStatus()),
			IdTransaction: transaction.GetInternalId(),
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

func (repo *StatusTransactionsGorm) Update(entity ent.StatusTransaction, conn repo.IConnection) error {

	ent := entities.StatusTransaction{
		Uuid:      entity.GetID(),
		Status:    uint8(entity.GetStatus()),
		UpdatedAt: entity.UpdatedAt,
	}

	ommitedFields := []string{"IdTransaction", "CreatedAt", "DeletedAt"}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		conn2.Tx.Omit(ommitedFields...).Save(&ent)
	} else {
		gorm.Connection.Db.Omit(ommitedFields...).Save(&ent)
	}

	return nil
}

func (repo *StatusTransactionsGorm) Delete(entity ent.StatusTransaction, conn repo.IConnection) error {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Delete(&entities.StatusTransaction{}, entity.GetID()); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *StatusTransactionsGorm) DeleteMany(entitiesToDelete []ent.StatusTransaction, conn repo.IConnection) error {
	var uuidsToDelete []uuid.UUID

	for _, item := range entitiesToDelete {
		uuidsToDelete = append(uuidsToDelete, item.GetID())
	}

	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Where("uuid IN(?)", uuidsToDelete).Delete(&entities.StatusTransaction{}); result.Error != nil {
		return result.Error
	}

	return nil
}

func (repo *StatusTransactionsGorm) DeleteByTransaction(entity ent.Transaction, conn repo.IConnection) error {
	if conn2, ok := conn.(*gorm.ConnectionGorm); ok && conn2 != nil {
		repo.db = conn2.Tx
	} else {
		repo.db = gorm.Connection.Db
	}

	if result := repo.db.Raw(
		"UPDATE accounting.status_transactions st "+
			"set deleted_at = now() FROM accounting.transactions t "+
			"where t.id = st.id_transaction and t.uuid =?",
		entity.GetID(),
	).Scan(nil); result.Error != nil {
		return result.Error
	}

	return nil
}

func NewStatusTransactionsGorm() repositories.IStatusTransactionsRepository {
	return &StatusTransactionsGorm{}
}
