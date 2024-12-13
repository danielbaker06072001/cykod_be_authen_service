package Interface

import (
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Model"
)

type IUserData interface {
	GetUserByUsername(username string) (*Model.UserProfile, error)
	GetUserByEmail(email string) (*Model.UserProfile, error)
	CheckUserExists(username string, email string) (bool, error)
	CreateUser(params DTO.Param) (*Model.UserProfile, error)
}