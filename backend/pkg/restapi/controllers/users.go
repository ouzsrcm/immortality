package restapi

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
)

type AuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// / Index godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} bool "true"
// @Router /users [get]
func UserList(w http.ResponseWriter, _ *http.Request) {

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
// @Tags auth
// @Accept  json
// @Produce  json
// @Param request body AuthDto true "request"
// @Success 200 {object} bool "true"
// @Router /auth [post]
func Auth(w http.ResponseWriter, r *http.Request) {

	var model AuthDto

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&model)

	userStore := users.UserStore{}
	userStore.Connect()

	res, _ := userStore.VerifyCredential(model.Email, model.Password)
	user, _ := userStore.GetUserByEmail(model.Email)

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
