package users

import (
	"errors"

	"immortality/pkg/common"

	"gorm.io/gorm"
)

const USERS = "users"
const CREDENTIALS = "credentials"
const CREDENTIALCHANGES = "credentialchanges"

type UserStore struct {
	common.StoreBase
}

type IUserStore interface {

	// User methods
	GetUser(id uint) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error
	UserIfExistsByEmail(email string) (bool, error)
	UserIfExistsByGsm(gsm string) (bool, error)

	// Credentials methods
	VerifyCredential(email string, password string) (*User, error)
	GetCredentials(id uint) (*Credential, error)
	GetCredentialsByUsername(username string) (*Credential, error)
	GetCredentialsByUserId(userId string) (*Credential, error)
	CreateCredentials(credentials *Credential) error
	UpdateCredentials(credentials *Credential) error
	DeleteCredentials(id string) error
}

func (s *UserStore) VerifyCredential(email string, password string) (bool, error) {

	var user *User
	var res *gorm.DB

	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res = tx.Where("email = ?", email).Table(USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		var credentials *Credential
		res = tx.Where("user_id = ?", user.ID).Table(CREDENTIALS).First(&credentials)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("credentials not found")
		}
		if credentials.Password != password {
			return errors.New("wrong password")
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}

func (s *UserStore) GetUser(id uint) (*User, error) {
	var user *User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res := tx.First(&User{}, id).Table(USERS)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("User not found")
		}
		res.First(&user)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return user, nil
}

func (s *UserStore) UserIfExistsByEmail(email string) (bool, error) {
	var user *User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("email = ?", email).Table(USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}

func (s *UserStore) UserIfExistsByGsm(gsm string) (bool, error) {
	var user *User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("gsm = ?", gsm).Table(USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}

func (s *UserStore) CreateUser(model *User) (*User, error) {
	var user *User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&model)
		if res.Error != nil {
			return res.Error
		}
		user, _ = s.GetUser(model.ID)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return user, nil
}
