package controllers

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	userStore := users.UserStore{}
	userStore.Connect()
	res, _ := userStore.GetUsers()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(res)
	w.Write([]byte(json))

}
