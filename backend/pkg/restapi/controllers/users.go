package restapi

import (
	"encoding/json"
	"immortality/pkg/common"
	"immortality/pkg/restapi/apibase"
	"immortality/pkg/users"
	"net/http"
	"time"
)

type UserDto struct {
	ID        uint   `json:"id" example:"1"`
	Email     string `json:"email" example:"john.doe@gmail.com"`
	Gsm       string `json:"gsm" example:"555-555-5555"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
}

type CredentialDto struct {
	common.ApiResponse

	ID       uint   `json:"id" example:"1"`
	UserId   uint   `json:"userId" example:"1"`
	Email    string `json:"email" example:"john.doe@gmail.com"`
	Password string `json:"password" example:"123456"`
}

type UserListResponse struct {
	common.ApiResponse

	Users []UserDto `json:"users"`
}

type UserCreateRequest struct {
	Email     string `json:"email" example:"john.doe@gmail.com"`
	Password  string `json:"password" example:"123456"`
	Gsm       string `json:"gsm" example:"555-555-5555"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
}

type UserCreateResponse struct {
	common.ApiResponse

	User       UserDto       `json:"user"`
	Credential CredentialDto `json:"credential"`
}

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
func UserList(w http.ResponseWriter, _ *http.Request) {

	userStore := users.NewUserStore()
	res, err := userStore.GetUsers()

	var users []UserDto
	var response UserListResponse

	w.Header().Set("Content-Type", "application/json")

	if err != nil {

		w.WriteHeader(http.StatusInternalServerError)

		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		json, _ := json.Marshal(response)

		w.Write([]byte(json))

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
	response.ErrorMessage = ""

	w.WriteHeader(http.StatusOK)

	json, _ := json.Marshal(response)
	w.Write([]byte(json))

}

// / Create godoc
// @Summary Create a user
// @Description Create a user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body UserCreateRequest true "User"
// @Success 200 {object} UserCreateResponse
// @Failure 400 {object} UserCreateResponse
// @Failure 500 {object} UserCreateResponse
// @Router /users [post]
func UserCreate(w http.ResponseWriter, r *http.Request) {

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
		w.WriteHeader(http.StatusInternalServerError)

		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		json, _ := json.Marshal(response)

		w.Write([]byte(json))

		return
	}

	credential.UserId = res.ID
	credential.Email = request.Email
	credential.Password = request.Password
	credential.CreatedAt = time.Now()

	creres, err := userStore.CreateCredential(&credential)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		response.Status = common.ApiStatusError
		response.ErrorMessage = err.Error()
		json, _ := json.Marshal(response)

		w.Write([]byte(json))

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
	response.ErrorMessage = ""

	w.WriteHeader(http.StatusOK)

	json, _ := json.Marshal(response)
	w.Write([]byte(json))

}
