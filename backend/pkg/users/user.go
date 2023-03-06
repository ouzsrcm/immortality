package users

import (
	"fmt"
	"immortality/pkg/database"
)

func SetupModule() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

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
	fmt.Println("Migrating database... ")
	return database.Migrate(dst...)
}
