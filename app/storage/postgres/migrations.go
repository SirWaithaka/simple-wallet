package postgres

import (
	"log"

	"simple-wallet/app/account"
	"simple-wallet/app/storage"
	"simple-wallet/app/transaction"
	"simple-wallet/app/user"
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
