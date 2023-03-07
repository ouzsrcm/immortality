package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string     `json:"email"`
	Gsm           string     `json:"gsm"`
	FirstName     string     `json:"firstName"`
	LastName      string     `json:"lastName"`
	LastLoginDate *time.Time `json:"lastLoginDate"`
}

type Credential struct {
	gorm.Model
	UserId   uint   `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CredentialChange struct {
	gorm.Model
	UserId       uint       `json:"userId"`
	CredentialId string     `json:"credentialId"`
	Hash         string     `json:"hash"` // guid
	Url          string     `json:"url"`
	ProcessDate  *time.Time `json:"processDate"`
	OldPassword  string     `json:"oldPassword"`
	NewPassword  string     `json:"newPassword"`
	IsChanged    bool       `json:"isChanged"`
}

type UserToken struct {
	gorm.Model

	UserId         uint      `json:"userId"`
	Token          string    `json:"token"`
	ExpirationDate time.Time `json:"expirationDate"`
	IsActive       bool      `json:"isActive"`
}

func SeedUser(db *gorm.DB) error {

	db.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&User{
			Email:     "oguzhan.saricam@gmail.com",
			FirstName: "Oğuz",
			LastName:  "Sarıçam",
		}).Error; err != nil {
			return err
		}

		// if err := tx.Create(&Credential{
		// 	ID:       util.GetUUID(),
		// 	UserId:   userId,
		// 	Username: "ouzsrcm",
		// 	Email:    "oguzhan.saricam@gmail.com",
		// 	Password: "123456"}).Error; err != nil {
		// 	return err
		// }

		return nil

	})

	return nil
}
