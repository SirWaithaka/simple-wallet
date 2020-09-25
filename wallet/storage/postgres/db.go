package postgres

import (
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"wallet"
	"wallet/storage"
)

var once sync.Once

// NewDatabase creates a new Database object
func NewDatabase(config wallet.Config) (*storage.Database, error) {
	var err error

	//var db *storage.Database
	db := new(storage.Database)

	var conn *gorm.DB
	once.Do(func() {
		dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s", config.DB.User)
		conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})

	if err != nil {
		return nil, err
	}
	db.DB = conn

	return db, nil
}
