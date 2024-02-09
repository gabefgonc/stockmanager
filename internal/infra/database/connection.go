package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDatabaseConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db/dev.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
