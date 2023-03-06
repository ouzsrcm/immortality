package users

type UserStore interface {

	// User methods
	GetUser(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(user *User) error
	UpdateUser(user *User) error
	DeleteUser(id string) error

	UserIfExistsByEmail(email string) (bool, error)

	UserIfExistsByGsm(gsm string) (bool, error)

	// Credentials methods
	GetCredentials(id string) (*Credential, error)
	GetCredentialsByUsername(username string) (*Credential, error)
	GetCredentialsByUserId(userId string) (*Credential, error)
	CreateCredentials(credentials *Credential) error
	UpdateCredentials(credentials *Credential) error
	DeleteCredentials(id string) error
}

// Path: pkg\users\userstore.go

func init() {
}

func GetUser(id string) (*User, error) {

	return nil, nil
}

func UserIfExistsByEmail(email string) (bool, error) {
	return false, nil
}

func UserIfExistsByGsm(gsm string) (bool, error) {
	return false, nil
}

func CreateUser(user *User) error {
	return nil
}
