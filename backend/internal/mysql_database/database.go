package mysql_database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "root:saricamou2@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to database... ")
	return db, nil
}

func AutoMigrate(dst ...interface{}) error {
	db, err := Connect()
	if err != nil {
		panic(err)
	}
	return db.AutoMigrate(dst...)
}
