package restapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() {
	fmt.Println("Initializing restapi...")

	r := mux.NewRouter()

	Routes(r)

	// TODO: routes

	// port'u config'e taşı
	http.ListenAndServe(":8080", r)

}
