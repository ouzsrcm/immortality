package users

import (
	"errors"
	"immortality/pkg/database"

	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
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
	GetCredentials(id uint) (*Credential, error)
	GetCredentialsByUsername(username string) (*Credential, error)
	GetCredentialsByUserId(userId string) (*Credential, error)
	CreateCredentials(credentials *Credential) error
	UpdateCredentials(credentials *Credential) error
	DeleteCredentials(id string) error
}

func (s *UserStore) Connect() (*gorm.DB, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, err
	}
	s.db = db
	return db, nil
}

func (s *UserStore) GetUser(id uint) (*User, error) {
	res := s.db.First(&User{}, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("User not found")
	}
	var user *User
	res.First(&user)
	return user, nil
}

func (s *UserStore) UserIfExistsByEmail(email string) (bool, error) {
	var user *User
	res := s.db.Where("email = ?", email).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func (s *UserStore) UserIfExistsByGsm(gsm string) (bool, error) {
	var user *User
	res := s.db.Where("gsm = ?", gsm).First(&user)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return false, nil
	}
	return true, nil
}

func (s *UserStore) CreateUser(model *User) (*User, error) {
	res := s.db.Create(&model)
	if res.Error != nil {
		return nil, res.Error
	}
	user, _ := s.GetUser(model.ID)
	return user, nil
}
