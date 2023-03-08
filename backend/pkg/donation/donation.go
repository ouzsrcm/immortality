package donation

import (
	"immortality/pkg/database"
)

func Setup() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

	// SeedingDatabase()
}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &DonationGroup{}, &Donation{}, &DonationAmount{}, &DonationMedia{}, &Media{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
