package Service

import (
	"fmt"
	"wan-api-verify-user/DTO"
	DTORegister "wan-api-verify-user/DTO/RegisterDTO"
	"wan-api-verify-user/Model"
	Interface "wan-api-verify-user/Service/User/Interafce"
	"wan-api-verify-user/Utils"
)

type UserService struct {
	UserDL Interface.IUserData
}

func NewUserServiceLayer(UserDL Interface.IUserData) Interface.IUserService {
	return &UserService{
		UserDL: UserDL,
	}
}

func (UserService *UserService) GetUserByUsername(username string) (*Model.Userprofile, error) {
	panic("implement me")
}

func (UserService *UserService) RegisterUser(params DTO.Param) (*DTORegister.RegisterInputDTO, error) {
	var registerDTO DTORegister.RegisterInputDTO
	
	// Do A check if the user is already registered
	username := Utils.ConvertInterface(params["Username"])
	email := Utils.ConvertInterface(params["Email"])

	userprofileModel, err := UserService.UserDL.CheckUserExists(username, email)
	if err == nil {
		return nil, fmt.Errorf("user already exists with username: %s or email: %s", username, email)
	}

	// Create a new user if the user does not exist
	
	fmt.Print(userprofileModel)
	return &registerDTO, nil
}
