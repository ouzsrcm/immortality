package routes

import (
	_ "immortality/docs"
	restapi "immortality/pkg/restapi/controllers"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {

	usersRoute := r.PathPrefix("/users").Subrouter()
	usersRoute.HandleFunc("", restapi.UserList).Methods("GET")
	usersRoute.HandleFunc("/get/{id}", restapi.Get).Methods("GET")
	usersRoute.HandleFunc("/update/{id}", restapi.Update).Methods("PUT")
	usersRoute.HandleFunc("/create", restapi.Create).Methods("POST")

	authRoute := r.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("", restapi.Auth).Methods("POST")
	authRoute.HandleFunc("/expire_token", restapi.ExpireToken).Methods("POST")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
