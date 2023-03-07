package controllers

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
)

// / Index godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} bool "true"
// @Router /users [get]
func Index(w http.ResponseWriter, _ *http.Request) {

	userStore := users.UserStore{}
	userStore.Connect()
	res, _ := userStore.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(res)
	w.Write([]byte(json))

}

// / Register godoc
// @Summary auth for token
// @Description auth for token
// @Tags users
// @Accept  json
// @Produce  json
// @Param username formData string true "Username"
// @Param password formData string true "Password"
// @Success 200 {object} bool "true"
// @Router /auth [post]
func Auth(w http.ResponseWriter, r *http.Request) {

	userStore := users.UserStore{}
	userStore.Connect()

	email := r.FormValue("email")
	password := r.FormValue("password")

	res, _ := userStore.VerifyCredential(email, password)
	user, _ := userStore.GetUserByEmail(email)

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
