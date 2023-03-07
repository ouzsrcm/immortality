package initializer

import (
	"immortality/pkg/benefactors"
	"immortality/pkg/users"
)

func Initialize() {

	users.Setup()

	benefactors.Setup()

	// TODO: setup other modules

}
