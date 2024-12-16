package Service

import (
	"fmt"
	"sync"
	"wan-api-verify-user/DTO"
	DTOLogin "wan-api-verify-user/DTO/LoginDTO"
	DTORegister "wan-api-verify-user/DTO/RegisterDTO"
	"wan-api-verify-user/Model"
	Interface "wan-api-verify-user/Service/User/Interafce"
	"wan-api-verify-user/Utils"

	"golang.org/x/crypto/bcrypt"
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

func (UserService *UserService) LoginUser(params DTO.Param) (*DTOLogin.LoginOutputDTO, error) {
	var LoginOutputDTO DTOLogin.LoginOutputDTO

	username := Utils.ConvertInterface(params["Username"])
	password := Utils.ConvertInterface(params["Password"])
	
	// Step 2.1: Check if the user exist in system by Redis, if not exist then return error
	// We only check if username exist, email is not used for login
	err := UserService.UserDL.CheckUserExists(username, "")
	if err == nil {
		return nil, fmt.Errorf("user does not exist, please register")
	}

	// Step 2.2: If user exist in redis, then continue check Redis if this user currently active
	// TODO : considering save name, email and role inside redis
	err, passhash, salt := UserService.UserDL.CheckUserExistsActive(username)
	if err == nil { // User is active return success message, (LOGIN SUCCESS)
		// Step 2.2.2: Check if the password is matched with the passhash and salt
		err = bcrypt.CompareHashAndPassword([]byte(passhash), []byte(password + salt))
		if err != nil {
			return nil, fmt.Errorf("password is incorrect")
		}

		// Generate the access Token and Refresh Token using go routine
		AccessToken, RefreshToken, err := UserService.GenerateToken(&Model.UserProfile{Username: username})
		if err != nil {
			return nil, err
		}

		LoginOutputDTO.Username = username
		LoginOutputDTO.Token.AccessToken = AccessToken
		LoginOutputDTO.Token.RefreshToken = RefreshToken
		return &LoginOutputDTO, err
	} 

	// Step 2.3: If user is not active, then check if the user is exist in database
	userprofileModel, err := UserService.UserDL.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	// Step 2.3.2: Check if the password is matched with the passhash and salt from the database
	err = bcrypt.CompareHashAndPassword([]byte(userprofileModel.Password), []byte(password + userprofileModel.Salt))
	if err != nil {
		return nil, fmt.Errorf("password is incorrect")
	}
	// Step 2.3.1: If user is exist in database, then return success message and set the user to Redis (active)
	LoginOutputDTO.Username = userprofileModel.Username
	LoginOutputDTO.Email = userprofileModel.Email
	LoginOutputDTO.FirstName = userprofileModel.FirstName
	LoginOutputDTO.LastName = userprofileModel.LastName

	// Generate the access Token and Refresh Token using go routine
	AccessToken, RefreshToken, err := UserService.GenerateToken(userprofileModel)
	if err != nil {
		return nil, err
	}

	LoginOutputDTO.Token.AccessToken = AccessToken
	LoginOutputDTO.Token.RefreshToken = RefreshToken

	// Set the user to Redis (active), this is to prevent the user to login again using postgre sql, more efficient
	err = UserService.UserDL.SetUserActive(username, userprofileModel.Password, userprofileModel.Salt)
	if err != nil {
		return nil, err
	}

	return &LoginOutputDTO, nil
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

func (UserService *UserService) GenerateToken(userprofileModel *Model.UserProfile) (string, string, error) {
	var claimsAccess, claimsRefresh DTO.DicodeClaims

	generateAccessToken := make(chan string)
	generateRefreshToken := make(chan string)
	generateError := make(chan error, 2) // Buffer of 2 to handle potential errors from both goroutines

	go func() {
		claimsAccess.Username = userprofileModel.Username
		claimsAccess.Roles = []string{"test_user"}
		AccessToken, err := Utils.GenerateJWT(claimsAccess, "ACCESS")
		if err != nil {
			generateError <- err
			return
		}
		generateAccessToken <- AccessToken
		close(generateAccessToken)
	}()

	go func() {
		claimsRefresh.Username = userprofileModel.Username
		RefreshToken, err := Utils.GenerateJWT(claimsRefresh, "REFRESH")
		if err != nil {
			generateError <- err
			return
		}
		generateRefreshToken <- RefreshToken
		close(generateRefreshToken)
	}()

	var AccessToken, RefreshToken string
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		select {
		case err := <-generateError:
			fmt.Println("Error generating access token:", err)
			return
		case AccessToken = <-generateAccessToken:
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case err := <-generateError:
			fmt.Println("Error generating refresh token:", err)
			return
		case RefreshToken = <-generateRefreshToken:
		}
	}()

	wg.Wait()

	// Ensure the error channel is read in case there are any remaining errors
	select {
	case err := <-generateError:
		return "", "", err
	default:
	}

	if AccessToken == "" || RefreshToken == "" {
		return "", "", fmt.Errorf("failed to generate tokens")
	}

	return AccessToken, RefreshToken, nil
}