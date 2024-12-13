package Interface

import (
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Model"
)

type IUserData interface {
	GetUserByUsername(username string) (*Model.Userprofile, error)
	GetUserByEmail(email string) (*Model.Userprofile, error)
	CheckUserExists(username string, email string) (bool, error)
	CreateUser(params DTO.Param) (*Model.Userprofile, error)
}