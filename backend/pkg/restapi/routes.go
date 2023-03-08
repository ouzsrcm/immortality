package restapi

import (
	"fmt"
	_ "immortality/docs"
	restapi "immortality/pkg/restapi/controllers"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {

	usersRoute := r.PathPrefix("/users").Subrouter()
	usersRoute.HandleFunc("", restapi.UserList).Methods("GET")

	// usersRoute.HandleFunc("/{id}", restapi.UserDetail).Methods("GET")
	// usersRoute.HandleFunc("", restapi.UserCreate).Methods("POST")
	// usersRoute.HandleFunc("/{id}", restapi.UserUpdate).Methods("PUT")
	// usersRoute.HandleFunc("/{id}", restapi.UserDelete).Methods("DELETE")

	authRoute := r.PathPrefix("/auth").Subrouter()
	authRoute.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Auth: ", r.Body)

		restapi.Auth(w, r)

	}).Methods("POST")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}
