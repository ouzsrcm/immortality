package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email         string     `json:"email" example:"john.doe@gmail.com"`
	Gsm           string     `json:"gsm" example:"555-555-5555"`
	FirstName     string     `json:"firstName" example:"John"`
	LastName      string     `json:"lastName" example:"Doe"`
	LastLoginDate *time.Time `json:"lastLoginDate" example:"2021-01-01T00:00:00Z"`

	// Relations

	// UserToken
	UserToken []UserToken `json:"userToken" gorm:"foreignKey:UserId"`

	// Credential
	Credential []Credential `json:"credential" gorm:"foreignKey:UserId"`
}

type Credential struct {
	gorm.Model
	UserId   uint   `json:"userId" example:"1"`
	Username string `json:"username" example:"ouzsrcm"`
	Email    string `json:"email" example:"john.doe@gmail.com"`
	Password string `json:"password" example:"123456"`

	// Relations
	// User
	User User `json:"user" gorm:"foreignKey:UserId"`
}

type CredentialChange struct {
	gorm.Model
	UserId       uint       `json:"userId" example:"1"`
	CredentialId string     `json:"credentialId" example:"1"`
	Hash         string     `json:"hash" example:"UUID"`
	Url          string     `json:"url" example:"1"`
	ProcessDate  *time.Time `json:"processDate" example:"2021-01-01T00:00:00Z"`
	OldPassword  string     `json:"oldPassword" example:"123456"`
	NewPassword  string     `json:"newPassword" example:"1234567"`
	IsChanged    bool       `json:"isChanged" example:"true"`

	// Relations
	// Credential
	Credential Credential `json:"credential" gorm:"foreignKey:CredentialId"`
}

type UserToken struct {
	gorm.Model

	UserId         uint      `json:"userId" example:"1"`
	Token          string    `json:"token" example:"UUID"` // UUID
	ExpirationDate time.Time `json:"expirationDate" example:"2021-01-01T00:00:00Z"`
	IsActive       bool      `json:"isActive" example:"true"`
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

		if err := tx.Create(&Credential{
			UserId:   1,
			Username: "ouzsrcm",
			Email:    "oguzhan.saricam@gmail.com",
			Password: "123456"}).Error; err != nil {
			return err
		}

		return nil

	})

	return nil
}
