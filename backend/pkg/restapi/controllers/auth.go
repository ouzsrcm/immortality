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

	res, err := userStore.VerifyCredential(model.Email, model.Password)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	user, err := userStore.GetUserByEmail(model.Email)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
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
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
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

	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	resultInfo.StatusCode = http.StatusOK
	resultInfo.Data = response

	apibase.ApiResult(w, r, resultInfo)
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
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
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

	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	resultInfo.StatusCode = http.StatusOK
	resultInfo.Data = response

	apibase.ApiResult(w, r, resultInfo)
}

// / TokenExists godoc
// @Summary token exists
// @Description token exists
// @Tags auth
// @Accept json
// @Produce json
// @Param request body TokenExistsRequest true "request"
// @Success 200 {object} TokenExistsResponse "true"
// @Failure 401 {object} TokenExistsResponse "false"
// @Failure 404 {object} TokenExistsResponse "false"
// @Failure 500 {object} TokenExistsResponse "false"
// @Router /auth/token_exists [post]
func TokenExists(w http.ResponseWriter, r *http.Request) {
	var response TokenExistsResponse
	var model TokenExistsRequest
	json.NewDecoder(r.Body).Decode(&model)

	userStore := users.NewUserStore()

	res, err := userStore.TokenExists(model.Token)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	json, _ := json.Marshal(TokenExistsResponse{
		Exists: res,
	})
	response.Data = json
	response.Status = common.ApiStatusSuccess

	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	resultInfo.StatusCode = http.StatusOK
	resultInfo.Data = response

	apibase.ApiResult(w, r, resultInfo)
}

// / CurrentTokens godoc
// @Summary current tokens
// @Description current tokens
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} CurrentTokensResponse "true"
// @Failure 401 {object} CurrentTokensResponse "false"
// @Failure 404 {object} CurrentTokensResponse "false"
// @Failure 500 {object} CurrentTokensResponse "false"
// @Router /auth/current_tokens [get]
func CurrentTokens(w http.ResponseWriter, r *http.Request) {

	var response CurrentTokensResponse

	userStore := users.NewUserStore()

	res, err := userStore.GetTokens()
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	tokens := make(map[string]uint)
	for _, token := range res {
		tokens[token.Token] = token.UserId
	}
	response.Data = &CurrentTokensResponse{
		Tokens: tokens,
	}
	response.Status = common.ApiStatusSuccess

	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	resultInfo.StatusCode = http.StatusOK
	resultInfo.Data = response

	apibase.ApiResult(w, r, resultInfo)
}

// / ExpireAllTokens godoc
// @Summary expire all tokens
// @Description expire all tokens
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} ExpireAllTokensResponse "true"
// @Failure 401 {object} ExpireAllTokensResponse "false"
// @Failure 404 {object} ExpireAllTokensResponse "false"
// @Failure 500 {object} ExpireAllTokensResponse "false"
// @Router /auth/expire_all_tokens [get]
func ExpireAllTokens(w http.ResponseWriter, r *http.Request) {

	var response ExpireAllTokensResponse

	userStore := users.NewUserStore()

	res, err := userStore.ExpireAllTokens()
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusNotFound, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	response.Data = ExpireAllTokensResponse{}
	if res == true {
		response.Status = common.ApiStatusSuccess
	} else {
		response.Status = common.ApiStatusError
	}
	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	resultInfo.StatusCode = http.StatusOK
	resultInfo.Data = response
	apibase.ApiResult(w, r, resultInfo)
}
