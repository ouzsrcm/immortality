package main

import (
	"immortality/pkg/initializer"
	"immortality/pkg/restapi"
)

func main() {

	initializer.Initialize()

	restapi.Initialize()

}
