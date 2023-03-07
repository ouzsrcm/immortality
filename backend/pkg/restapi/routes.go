package restapi

import (
	"immortality/pkg/restapi/controllers"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/users", controllers.Index).Methods("GET")
}
