package restapi

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
	"time"
)

type UserListResponse struct {
	Email         string     `json:"email" example:"john.doe@gmail.com"`
	Gsm           string     `json:"gsm" example:"555-555-5555"`
	FirstName     string     `json:"firstName" example:"John"`
	LastName      string     `json:"lastName" example:"Doe"`
	LastLoginDate *time.Time `json:"lastLoginDate" example:"2021-01-01T00:00:00Z"`
}

// / Index godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} UserListResponse
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
