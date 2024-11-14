package Controller

import (
	"net/http"
	Interface "wan-api-verify-user/Service/User/Interafce"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService Interface.IUserService
}

func NewUserControllerLayer(context *gin.Engine, UserService Interface.IUserService) {
	UserControllerObject :=  &UserController{
		UserService: UserService,
	}

	context.GET("/user", func(c *gin.Context) { 
		c.JSON(http.StatusOK, map[string]string{"status": "healthy api"} )
	})

	AuthenticationGroup := context.Group("/authentication")
	 { 
		AuthenticationGroup.POST("/login", func(c *gin.Context) {
			UserControllerObject.login(c)
		})

		AuthenticationGroup.POST("/register", func(c *gin.Context) {
			UserControllerObject.register(c)
		})

		AuthenticationGroup.POST("/forgot-password", func(c *gin.Context) {
			UserControllerObject.forgotPassword(c)
		})
	}
}

func (UserController *UserController) login(context *gin.Context) {
	panic("implement me")
}

func (UserController *UserController) register(context *gin.Context) {
	panic("implement me")
}

func (UserController *UserController) forgotPassword(context *gin.Context) {
	panic("implement me")
}