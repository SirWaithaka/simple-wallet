package postgres

import (
	"log"
	"wallet/account"
	"wallet/storage"
	"wallet/transaction"
	"wallet/user"
)

// Migrate updates the db with new columns, and tables
func Migrate(database *storage.Database) {
	err := database.DB.AutoMigrate(
		user.User{},
		account.Account{},
		transaction.Transaction{},
	)

	if err != nil {
		log.Println(err)
	}
}
