package postgres

import (
	"wallet/account"
	"wallet/storage"
	"wallet/transaction"
	"wallet/user"
)

// Migrate updates the db with new columns, and tables
func Migrate(database *storage.Database) {
	database.DB.AutoMigrate(
		user.User{},
		account.Account{},
		transaction.Transaction{},
	)
}
