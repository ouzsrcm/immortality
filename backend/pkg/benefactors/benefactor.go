package benefactors

import (
	"immortality/pkg/database"
)

func Setup() {

	dst := GetInterfaces()

	MigrateDatabase(dst...)

}

func GetInterfaces() []interface{} {
	models := make([]interface{}, 0)
	return append(models, &Benefactor{}, &BenefactorAddress{}, &Company{}, &BenefactorCreditCard{}, &BenefactorBankAccount{})
}

func MigrateDatabase(dst ...interface{}) error {
	if dst != nil {
		dst = append(dst, GetInterfaces()...)
	}
	res := database.Migrate(dst...)
	return res
}
