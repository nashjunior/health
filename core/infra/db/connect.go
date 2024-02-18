package db

import "health/core/infra/db/gorm"

func StartConnections() {
	gorm.NewConnection()

	gorm.Connection.Connect()
}
