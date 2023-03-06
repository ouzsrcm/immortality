package users

import (
	"immortality/internal/database"
	"immortality/util"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Credential struct {
	gorm.Model
	ID       string `json:"id"`
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CredentialChange struct {
	gorm.Model
	ID           string     `json:"id"`
	UserId       string     `json:"userId"`
	CredentialId string     `json:"credentialId"`
	Hash         string     `json:"hash"` // guid
	Url          string     `json:"url"`
	ProcessDate  *time.Time `json:"processDate"`
	OldPassword  string     `json:"oldPassword"`
	NewPassword  string     `json:"newPassword"`
	IsChanged    bool       `json:"isChanged"`
}

func SeedUser(db *gorm.DB) error {

	userId := util.GetUUID()

	db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&User{
			ID:        userId,
			Email:     "oguzhan.saricam@gmail.com",
			FirstName: "Oğuz",
			LastName:  "Sarıçam",
		}).Error; err != nil {
			return err
		}

		if err := tx.Create(&Credential{
			ID:       util.GetUUID(),
			UserId:   userId,
			Username: "ouzsrcm",
			Email:    "oguzhan.saricam@gmail.com",
			Password: "123456"}).Error; err != nil {
			return err
		}

		return nil

	})

	return nil
}

func MigrateDatabase() error {
	return database.AutoMigrate(&User{}, &Credential{})
}
