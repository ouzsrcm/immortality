package restapi

import (
	"immortality/pkg/common"
)

type UserDto struct {
	ID        uint   `json:"id" example:"1"`
	Email     string `json:"email" example:"john.doe@gmail.com"`
	Gsm       string `json:"gsm" example:"555-555-5555"`
	FirstName string `json:"firstName" example:"John"`
	LastName  string `json:"lastName" example:"Doe"`
}

type UserResponse struct {
	common.ApiResponse
	User UserDto `json:"user"`
}

type UserUpdateRequest struct {
	UserDto
}

type UserUpdateResponse struct {
	common.ApiResponse
	User UserDto `json:"user"`
}

type UserDeleteRequest struct {
	ID uint `json:"id" example:"1"`
}
type UserDeleteResponse struct {
	common.ApiResponse
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
