package common

import (
	"immortality/pkg/database"

	"gorm.io/gorm"
)

type StoreBase struct {
	Db *gorm.DB
}

func (s *StoreBase) Connect() (*gorm.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	s.Db = db
	return db, nil
}
