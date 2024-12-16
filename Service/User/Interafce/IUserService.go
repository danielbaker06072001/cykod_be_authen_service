package Interface

import (
	"wan-api-verify-user/DTO"
	DTOLogin "wan-api-verify-user/DTO/LoginDTO"
	DTORegister "wan-api-verify-user/DTO/RegisterDTO"
	"wan-api-verify-user/Model"
)

type IUserService interface {
	GetUserByUsername(username string) (*Model.UserProfile, error)
	RegisterUser(user DTO.Param) (*DTORegister.RegisterInputDTO, error)
	LoginUser(user DTO.Param) (*DTOLogin.LoginOutputDTO, error)
	GenerateToken(user *Model.UserProfile) (string, string, error)
}