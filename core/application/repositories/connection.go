package repositories

import "database/sql"

type IConnection interface {
	Connect() error
	StartTransaction(opts ...*sql.TxOptions)
	CommitTransaction()
	RollbackTransaction()
}
