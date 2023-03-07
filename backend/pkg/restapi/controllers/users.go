package controllers

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	userStore := users.UserStore{}
	res, _ := userStore.GetUsers()

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		panic(err)
	}

}
