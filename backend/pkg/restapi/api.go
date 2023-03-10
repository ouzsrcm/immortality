package restapi

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() {
	fmt.Println("Initializing restapi...")

	r := mux.NewRouter()

	amw := AuthMiddleWare{}
	amw.Init()
	r.Use(amw.Middleware)

	Routes(r)

	// port'u config'e taşı
	http.ListenAndServe(":8080", r)
}
