package users

import "immortality/pkg/database"

func SetupModule() {

	MigrateDatabase()

	// seed data
	// TODO: other operations
}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &User{}, &Credential{}, &CredentialChange{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	return database.Migrate(dst...)
}
