package postgres

import (
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

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
		conn, err = gorm.Open("postgres", config.DB.String("disable"))
	})

	if err != nil {
		return nil, err
	}
	db.DB = conn

	return db, nil
}

