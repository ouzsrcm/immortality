package database

import (
	"immortality/internal/mysql_database"

	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return mysql_database.Connect()
}

func Migrate(dst ...interface{}) error {
	return mysql_database.AutoMigrate(dst...)
}
