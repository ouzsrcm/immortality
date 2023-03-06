package users

import (
	"fmt"
	"immortality/pkg/database"
)

func SetupModule() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	db, err := database.Connect()
	if err != nil {
		panic(err)
	}
	fmt.Println("Seeding users...")

	SeedUser(db)

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
	res := database.Migrate(dst...)
	return res
}
