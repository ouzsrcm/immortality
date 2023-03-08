package initializer

import (
	"immortality/pkg/benefactors"
	"immortality/pkg/donation"
	"immortality/pkg/order"
	"immortality/pkg/users"
)

func Initialize() {

	users.Setup()

	benefactors.Setup()

	donation.Setup()

	order.Setup()

}
