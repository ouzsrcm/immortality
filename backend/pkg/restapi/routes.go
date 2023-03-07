package restapi

import (
	"immortality/pkg/restapi/controllers"

	_ "immortality/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/users", controllers.Index).Methods("GET")
	r.HandleFunc("/auth", controllers.Auth).Methods("POST")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

}
