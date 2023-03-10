package restapi

import (
	"encoding/json"
	"immortality/pkg/common"
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/users"
	"immortality/util"
	"net/http"
	"time"
)

// TODO: apibase.NewResultInfo metodu hatalı.
/// düzeltilmesi gerekiyor.

// / Index godoc
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} UserListResponse
// @Failure 400 {object} UserListResponse
// @Failure 500 {object} UserListResponse
// @Router /users [get]
func UserList(w http.ResponseWriter, r *http.Request) {

	userStore := users.NewUserStore()
	res, err := userStore.GetUsers()

	var users []UserDto
	var response UserListResponse

	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	for _, user := range res {
		users = append(users, UserDto{
			ID:        user.ID,
			Email:     user.Email,
			Gsm:       user.Gsm,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		})
	}
	response.Users = users
	response.Status = common.ApiStatusSuccess
	resultInfo := apibase.NewResultInfo(http.StatusOK, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
}

// / Create godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body UserCreateRequest true "User"
// @Success 200 {object} UserResponse
// @Failure 400 {object} UserResponse
// @Failure 500 {object} UserResponse
// @Router /users [post]
func Create(w http.ResponseWriter, r *http.Request) {

	var user users.User
	var credential users.Credential
	var request UserCreateRequest
	var response UserCreateResponse

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}

	user.Email = request.Email
	user.Gsm = request.Gsm
	user.FirstName = request.FirstName
	user.LastName = request.LastName

	userStore := users.NewUserStore()

	res, err := userStore.CreateUser(&user)

	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}

	credential.UserId = res.ID
	credential.Email = request.Email
	credential.Password = request.Password
	credential.CreatedAt = time.Now()

	creres, err := userStore.CreateCredential(&credential)

	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		var resultInfo apibase.ResultInfo
		resultInfo.ContentType = "application/json"
		resultInfo.StatusCode = http.StatusBadRequest
		resultInfo.Data = response
		apibase.ApiResult(w, r, resultInfo)
		return
	}

	response.Credential = CredentialDto{
		ID:     creres.ID,
		UserId: creres.UserId,
		Email:  creres.Email,
		//Password:  creres.Password
	}

	response.User = UserDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = common.ApiStatusSuccess
	resultInfo := apibase.NewResultInfo(http.StatusOK, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
}

// / / Get godoc
// @Summary Get a user
// @Description Get a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} UserResponse
// @Failure 400 {object} UserResponse
// @Failure 500 {object} UserResponse
// @Router /users/{id} [get]
func Get(w http.ResponseWriter, r *http.Request) {
	response := UserResponse{}
	strId := r.URL.Query().Get("id")
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	userStore := users.NewUserStore()
	res, err := userStore.GetUser(id)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	response.User = UserDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = common.ApiStatusSuccess
	resultInfo := apibase.NewResultInfo(http.StatusOK, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
}

// / Update godoc
// @Summary Update a user
// @Description Update a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Param user body UserUpdateRequest true "User"
// @Success 200 {object} UserUpdateResponse
// @Failure 400 {object} UserUpdateResponse
// @Failure 500 {object} UserUpdateResponse
// @Router /users/{id} [put]
func Update(w http.ResponseWriter, r *http.Request) {
	var user UserUpdateRequest
	response := UserUpdateResponse{}
	strId := r.URL.Query().Get("id")
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	userStore := users.NewUserStore()
	res, _ := userStore.GetUser(id)

	res.Email = user.Email
	res.Gsm = user.Gsm
	res.FirstName = user.FirstName
	res.LastName = user.LastName

	res, err = userStore.UpdateUser(res)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	response.User = UserDto{
		ID:        res.ID,
		Email:     res.Email,
		Gsm:       res.Gsm,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	response.Status = common.ApiStatusSuccess
	resultInfo := apibase.NewResultInfo(http.StatusOK, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
}

// / Delete godoc
// @Summary Delete a user
// @Description Delete a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param id path uint true "User ID"
// @Success 200 {object} UserDeleteResponse
// @Failure 400 {object} UserDeleteResponse
// @Failure 500 {object} UserDeleteResponse
// @Router /users/{id} [delete]
func Delete(w http.ResponseWriter, r *http.Request) {
	response := UserDeleteResponse{}
	strId := r.URL.Query().Get("id")
	id, err := util.ToUint(strId)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	userStore := users.NewUserStore()
	err = userStore.DeleteUser(id)
	if err != nil {
		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		resultInfo := apibase.NewResultInfo(http.StatusBadRequest, err.Error(), "application/json", response)
		apibase.ApiResult(w, r, *resultInfo)
		return
	}
	response.Status = common.ApiStatusSuccess
	resultInfo := apibase.NewResultInfo(http.StatusOK, err.Error(), "application/json", response)
	apibase.ApiResult(w, r, *resultInfo)
}
