package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
)


// Database is a wrapper type for the gorm DB pointer
type Database struct {
	*gorm.DB
}

func(db *Database) Close() {
	err := db.DB.Close()
	if err != nil {
		fmt.Printf("Error closing db: %v\n", err)
	}
}