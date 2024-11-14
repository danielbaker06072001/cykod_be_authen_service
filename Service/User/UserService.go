package Service

import (
	"wan-api-verify-user/Model"
	Interface "wan-api-verify-user/Service/User/Interafce"
)

type UserService struct {
	UserDL Interface.IUserData
}

func NewUserServiceLayer(UserDL Interface.IUserData) Interface.IUserService {
	return &UserService{
		UserDL: UserDL,
	}
}

func (UserService *UserService) GetUserByUsername(username string) (*Model.KOL, error) {
	panic("implement me")
}
