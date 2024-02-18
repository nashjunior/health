package gorm

import (
	"database/sql"
	"health/core/application/repositories"
	"health/core/infra/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionGorm struct {
	Db *gorm.DB
	Tx *gorm.DB
}

func (transaction *ConnectionGorm) Connect() error {

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN: "host=" + config.PostgresHost +
					" user=" + config.PostgresUser +
					" dbname=" + config.PostgresDb +
					" password=" + config.PostgresUserPassword +
					" sslmode=disable",
			},
		),
	)

	if err != nil {
		return err
	}

	transaction.Db = db
	return nil
}

func (transaction *ConnectionGorm) StartTransaction(opts ...*sql.TxOptions) {
	transaction.Tx = transaction.Db.Begin(opts...)
}

func (transaction *ConnectionGorm) CommitTransaction() {
	transaction.Tx.Debug().Commit()
}
func (transaction *ConnectionGorm) RollbackTransaction() {
	transaction.Tx.Rollback()
}

var Connection ConnectionGorm

func NewConnection() repositories.IConnection {
	Connection = ConnectionGorm{}
	return &Connection
}
