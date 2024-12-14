package Service

import (
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

func (UserService *UserService) GetUserByUsername(username string) (*Model.UserProfile, error) {
	panic("implement me")
}

func (UserService *UserService) RegisterUser(params DTO.Param) (*DTORegister.RegisterInputDTO, error) {
	var registerDTO DTORegister.RegisterInputDTO
	
	// Do A check if the user is already registered
	username := Utils.ConvertInterface(params["Username"])
	email := Utils.ConvertInterface(params["Email"])


	// Check if the user already exists
	err := UserService.UserDL.CheckUserExists(username, email)
	if err != nil {
		return nil, err
	}

	// Hashing the password and applying the salt
	password := Utils.ConvertInterface(params["Password"])
	salt, err := Utils.GenerateSalt(10); // Generate a random salt
	if err != nil {
		return nil, err
	}
	hashedPassword, err := Utils.HashPassword(password, salt)
	if err != nil {
		return nil, err
	}
	
	// Store the salt in the params that later passed on to create user
	params["Salt"] = salt
	params["Password"] = hashedPassword

	// Create a new user if the user does not exist
	userprofileModel, err := UserService.UserDL.CreateUser(params)
	if err != nil {
		return nil, err
	}
	
	registerDTO.Email = userprofileModel.Email
	registerDTO.Username = userprofileModel.Username
	registerDTO.LastName = userprofileModel.LastName
	registerDTO.FirstName = userprofileModel.FirstName
	return &registerDTO, nil
}
