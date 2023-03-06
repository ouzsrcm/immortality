package database

import (
	"immortality/internal/sqlite_database"

	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return sqlite_database.Connect()
}

func Migrate(dst ...interface{}) error {
	return sqlite_database.AutoMigrate(dst...)
}
