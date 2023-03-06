package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	return db, err
}

func AutoMigrate(dst ...interface{}) error {
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	return db.AutoMigrate(dst...)
}
