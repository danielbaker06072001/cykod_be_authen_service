package Interface

import (
	"wan-api-verify-user/DTO"
	"wan-api-verify-user/Model"
)

type IUserData interface {
	GetUserByUsername(username string) (*Model.UserProfile, error)
	GetUserByEmail(email string) (*Model.UserProfile, error)
	CheckUserExists(username string, email string) (error)
	CheckUserExistsActive(username string) (error, string , string)

	SetUserActive(username string, passHash string, salt string) (error)
	CreateUser(params DTO.Param) (*Model.UserProfile, error)
}