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
		apibase.ApiResult(w, r, apibase.ResultInfo{
			StatusCode:  http.StatusNotFound,
			ContentType: "application/json",
			Data:        response,
		})
		return
	}
	user, err := userStore.GetUserByEmail(model.Email)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		apibase.ApiResult(w, r, apibase.ResultInfo{
			StatusCode:  http.StatusNotFound,
			ContentType: "application/json",
			Data:        response,
		})
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
		apibase.ApiResult(w, r, apibase.ResultInfo{
			StatusCode:  http.StatusBadRequest,
			ContentType: "application/json",
			Data:        response,
		})
		return
	}
	response.Data = AuthResponseData{
		Token:  token,
		UserId: user.ID,
	}
	apibase.ApiResult(w, r, apibase.ResultInfo{
		StatusCode:  http.StatusOK,
		ContentType: "application/json",
		Data:        response,
	})
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
		apibase.ApiResult(w, r, apibase.ResultInfo{
			ContentType: "application/json",
			StatusCode:  http.StatusNotFound,
			Data:        response,
		})
		return
	}
	if res != nil {
		response.Status = common.ApiStatusSuccess
		response.ErrorMessage = "Token Expired"
	} else {
		response.Status = common.ApiStatusError
		response.ErrorMessage = "Token not found"
	}
	apibase.ApiResult(w, r, apibase.ResultInfo{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
		Data:        response,
	})
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
	response.Exists = res
	response.Status = common.ApiStatusSuccess
	apibase.ApiResult(w, r, apibase.ResultInfo{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
		Data:        response,
	})
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
		apibase.ApiResult(w, r, apibase.ResultInfo{
			ContentType: "application/json",
			StatusCode:  http.StatusNotFound,
			Data:        response,
		})
		return
	}
	tokens := make(map[string]uint)
	for _, token := range res {
		tokens[token.Token] = token.UserId
	}
	response.Data = tokens
	response.Status = common.ApiStatusSuccess
	apibase.ApiResult(w, r, apibase.ResultInfo{
		ContentType: "application/json",
		StatusCode:  http.StatusOK,
		Data:        response,
	})
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
	userStore := users.NewUserStore()
	res, err := userStore.ExpireAllTokens()
	if err != nil {
		response := &ExpireAllTokensResponse{}
		response.Status = common.ApiStatusError
		apibase.ApiResult(w, r, apibase.ResultInfo{
			ContentType: "application/json",
			StatusCode:  http.StatusNotFound,
			Data:        response,
		})
		return
	}
	var resultInfo apibase.ResultInfo
	resultInfo.ContentType = "application/json"
	if res {
		resultInfo.StatusCode = http.StatusOK
	} else {
		resultInfo.StatusCode = http.StatusBadRequest
	}
	apibase.ApiResult(w, r, resultInfo)
}
