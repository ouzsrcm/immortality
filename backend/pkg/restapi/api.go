package restapi

import (
	"fmt"
	"immortality/pkg/restapi/middleware"
	"immortality/pkg/restapi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func Initialize() {
	fmt.Println("Initializing restapi...")

	r := mux.NewRouter()

	amw := middleware.AuthMiddleWare{}
	amw.Init()
	r.Use(amw.Middleware)

	routes.Routes(r)

	// port'u config'e taşı
	http.ListenAndServe(":8080", r)
}
