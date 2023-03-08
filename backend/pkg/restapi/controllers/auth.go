package restapi

import (
	"encoding/json"
	"immortality/pkg/users"
	"net/http"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	UserId       uint   `json:"user_id"`
	Email        string `json:"email"`
	Token        string `json:"token"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

type ExpireTokenRequest struct {
	Token string `json:"token"`
}

type ExpireTokenResponse struct {
	UserId       uint   `json:"user_id"`
	Token        string `json:"token"`
	StatusCode   int    `json:"status_code"`
	ErrorMessage string `json:"error_message"`
}

// / Register godoc
// @Summary auth for token
// @Description auth for token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body AuthRequest true "request"
// @Success 200 {object} AuthResponse "true"
// @Failure 401 {object} AuthResponse "false"
// @Failure 404 {object} AuthResponse "false"
// @Failure 500 {object} AuthResponse "false"
// @Router /auth [post]
func Auth(w http.ResponseWriter, r *http.Request) {

	var model AuthRequest
	var token string

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&model)

	userStore := users.UserStore{}
	userStore.Connect()

	res, _ := userStore.VerifyCredential(model.Email, model.Password)
	user, _ := userStore.GetUserByEmail(model.Email)

	w.Header().Set("Content-Type", "application/json")
	if res {
		tokenres, err := userStore.GenerateToken(user)
		if err != nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		token = tokenres.Token
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}

	json, _ := json.Marshal(AuthResponse{
		UserId:       user.ID,
		Email:        user.Email,
		Token:        token,
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
	})

	w.Write([]byte(json))
}

// / ExpireToken godoc
// @Summary expire token
// @Description expire token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body ExpireTokenRequest true "request"
// @Success 200 {object} ExpireTokenResponse "true"
// @Failure 401 {object} ExpireTokenResponse "false"
// @Failure 404 {object} ExpireTokenResponse "false"
// @Failure 500 {object} ExpireTokenResponse "false"
// @Router /auth/expire_token [post]
func ExpireToken(w http.ResponseWriter, r *http.Request) {
	var model ExpireTokenRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&model)

	userStore := users.UserStore{}
	userStore.Connect()

	w.Header().Set("Content-Type", "application/json")

	res, err := userStore.ExpireToken(model.Token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	json, _ := json.Marshal(ExpireTokenResponse{
		UserId:       res.UserId,
		Token:        res.Token,
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
	})

	w.Write([]byte(json))

}
