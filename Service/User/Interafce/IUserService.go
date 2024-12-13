package Interface

import (
	"wan-api-verify-user/DTO"
	DTORegister "wan-api-verify-user/DTO/RegisterDTO"
	"wan-api-verify-user/Model"
)

type IUserService interface {
	GetUserByUsername(username string) (*Model.UserProfile, error)
	RegisterUser(user DTO.Param) (*DTORegister.RegisterInputDTO, error)
}