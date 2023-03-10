// path: backend\pkg\restapi\restapi.go

package restapi

import (
	"encoding/json"
	"immortality/pkg/common"
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/users"
	"net/http"
)

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

	var response AuthResponse
	var model AuthRequest
	var token string

	json.NewDecoder(r.Body).Decode(&model)

	userStore := users.NewUserStore()

	res, _ := userStore.VerifyCredential(model.Email, model.Password)
	user, err := userStore.GetUserByEmail(model.Email)

	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	if res {
		tokenres, err := userStore.GenerateToken(user)
		if err != nil {
			response.Status = common.ApiStatusError
		} else {
			response.Status = common.ApiStatusSuccess
		}
		token = tokenres.Token
	} else {
		response.Status = common.ApiStatusError
		response.ErrorMessage = "Invalid email or password"
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	json, _ := json.Marshal(AuthResponse{
		UserId:       user.ID,
		Email:        user.Email,
		Token:        token,
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
	})
	response.Data = json
	resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
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

	var response ExpireTokenResponse
	var model ExpireTokenRequest
	json.NewDecoder(r.Body).Decode(&model)

	userStore := users.NewUserStore()

	res, err := userStore.ExpireToken(model.Token)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	json, _ := json.Marshal(ExpireTokenResponse{
		UserId:       res.UserId,
		Token:        res.Token,
		StatusCode:   http.StatusOK,
		ErrorMessage: "",
	})
	response.Data = json
	response.Status = common.ApiStatusSuccess
	resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
}
