package Controller

import (
	"net/http"
	"wan-api-verify-user/AppConfig/Common"
	DTOParam "wan-api-verify-user/DTO"
	DTOLogin "wan-api-verify-user/DTO/LoginDTO"
	DTO "wan-api-verify-user/DTO/RegisterDTO"
	Interface "wan-api-verify-user/Service/User/Interafce"
	ViewModel "wan-api-verify-user/ViewModel/UserprofileViewModel"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService Interface.IUserService
}

func NewUserControllerLayer(context *gin.Engine, UserService Interface.IUserService) {
	UserControllerObject :=  &UserController{
		UserService: UserService,
	}

	context.GET("/user/healthz", func(c *gin.Context) { 
		c.JSON(http.StatusOK, map[string]string{"status": "healthy api"} )
	})

	AuthenticationGroup := context.Group("/authentication")
	 { 
		AuthenticationGroup.POST("/login", func(c *gin.Context) {
			UserControllerObject.login(c)
		})

		AuthenticationGroup.POST("/logout", func(c *gin.Context) {
			panic("implement me")
		})

		AuthenticationGroup.POST("/register", func(c *gin.Context) {
			UserControllerObject.register(c)
		})

		AuthenticationGroup.POST("/forgot-password", func(c *gin.Context) {
			UserControllerObject.forgotPassword(c)
		})
	}
}

/*
	* This function login a user to the system based on the provided username and password
	Steps: 
		1. Check if all fields are filled
		2. Check if the user exist in the system using redis 
			2.1 First check if the user exist in Redis, if not exist then return error
			2.2 If user exist in redis , then continue check Redis if this user currently active
				2.2.1 If user is active, then get passhash and salt from Redis
				2.2.2 Check if the password is matched with the passhash and salt
				2.2.3 If password is matched, then return success message
			2.3 If user is not active, then check if the user is exist in database
				2.3.1 If user is exist in database, then return success message
					2.3.1.1 Then set the user to Redis (active)
				2.3.2 If user is not exist in database, then return error message
		3. Return a success message
*/
func (UserController *UserController) login(context *gin.Context) {
	var LoginOutputVM ViewModel.LoginViewModel

	var input DTOLogin.LoginInputDTO
	if err := context.BindJSON(&input); err != nil {
		LoginOutputVM.CommonResponse.Message = "all fields are required"
		context.JSON(http.StatusBadRequest, LoginOutputVM)
		return
	}

	// * Check if all fields are filled
	if input.Username == "" || input.Password == "" {
		LoginOutputVM.CommonResponse.Status = Common.FAIL
		LoginOutputVM.CommonResponse.Message = "all fields are required"
		context.JSON(http.StatusBadRequest, LoginOutputVM)
		return
	}

	var params DTOParam.Param = make(DTOParam.Param)
	params["Username"] = input.Username
	params["Password"] = input.Password

	LoginOutput, err := UserController.UserService.LoginUser(params)
	if err != nil {
		LoginOutputVM.CommonResponse.Status = Common.FAIL
		LoginOutputVM.CommonResponse.Message = err.Error()
		context.JSON(http.StatusBadRequest, LoginOutputVM)
		return
	}

	LoginOutputVM.CommonResponse.Status = Common.SUCCESS
	LoginOutputVM.CommonResponse.Message = "user logged in successfully"
	LoginOutputVM.Data = LoginOutput
	LoginOutputVM.Meta = &struct {
		NextStep []string `json:"next_steps"`
	}{}
	LoginOutputVM.Meta.NextStep = []string{"use the token to access the system"}

	context.JSON(http.StatusOK, LoginOutputVM)
}


/*
	* This function register a new user to the system
	Stepss:
		1. Check if all fields are filled
		2. Check if the user already exists using redis, if not create a new user
			2.1 If the User is not already exist, create user and also set new in Redis
		3. Return a success message
*/
func (UserController *UserController) register(context *gin.Context) {
	var RegisterVM ViewModel.RegisterViewModel

	var input DTO.RegisterInputDTO
	if err := context.BindJSON(&input); err != nil {
		RegisterVM.CommonResponse.Status = Common.FAIL
		context.JSON(http.StatusBadRequest, RegisterVM)
		return
	}

	// * Check if all fields are filled
	if input.Username == "" || input.Password == "" || input.Email == "" || input.FirstName == "" || input.LastName == "" {
		RegisterVM.CommonResponse.Status = Common.FAIL
		RegisterVM.CommonResponse.Message = "All fields are required"
		context.JSON(http.StatusBadRequest, RegisterVM)
		return
	}

	var params DTOParam.Param = make(DTOParam.Param)
	params["Username"] = input.Username
	params["Password"] = input.Password
	params["Email"] = input.Email
	params["FirstName"] = input.FirstName
	params["LastName"] = input.LastName

	RegisterDto , err := UserController.UserService.RegisterUser(params)
	if err != nil {
		RegisterVM.CommonResponse.Status = Common.FAIL
		RegisterVM.CommonResponse.Message = err.Error()
		context.JSON(http.StatusBadRequest, RegisterVM)
		return
	}

	RegisterVM.CommonResponse.Status = Common.SUCCESS
	RegisterVM.CommonResponse.Message = "User registered successfully"
	RegisterVM.Data = RegisterDto
	RegisterVM.Meta = &struct {
        NextStep []string `json:"next_steps"`
    }{}
	RegisterVM.Meta.NextStep = []string{"Verify your email to activate your account"}

	context.JSON(http.StatusOK, RegisterVM)
}

func (UserController *UserController) forgotPassword(context *gin.Context) {
	panic("implement me")
}