package controllers

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {

	userStore := users.UserStore{}
	userStore.Connect()
	res, _ := userStore.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(res)
	w.Write([]byte(json))

}

func Auth(w http.ResponseWriter, r *http.Request) {

	userStore := users.UserStore{}
	userStore.Connect()

	username := r.FormValue("username")
	password := r.FormValue("password")

	res, _ := userStore.VerifyCredential(username, password)
	user, _ := userStore.GetUserByEmail(username)

	w.Header().Set("Content-Type", "application/json")
	if res {
		tokenres, _ := userStore.GenerateToken(user)
		if tokenres {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
	json, _ := json.Marshal(res)
	w.Write([]byte(json))
}
