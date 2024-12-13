package main

import (
	"log"
	"net/http"
	"os"
	"wan-api-verify-user/AppConfig/Config"
	"wan-api-verify-user/Controller"
	"wan-api-verify-user/Data"
	Service "wan-api-verify-user/Service/KOL"
	UserService "wan-api-verify-user/Service/User"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// * Load environemnt file and start connection to database
	env := ".env"
	if len(os.Args) > 1 {
		env = os.Args[1]
	}

	Config.SetEnvironment(env)
	config, err := Config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db_data, err := Config.Connect(config)
	if err != nil {
		log.Fatal("Error connecting to database")
	}
	db_redis, err  := Config.ConnectRedis(config)
	if err != nil {
		log.Fatal("Error connecting to redis")
	}
	
	dataLayer := Data.NewKolDataLayer(db_data)
	service := Service.NewKOLService(dataLayer)
	Controller.NewKOLController(router, service)
	
	userDataLayer := Data.NewUserDataLayer(db_data, db_redis)
	userService := UserService.NewUserServiceLayer(userDataLayer)
	Controller.NewUserControllerLayer(router, userService)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "WRONG API PATH"})
	})

	if err := router.Run(":" + config.Server.GinPort); err != nil {
		log.Fatal("FAILED TO START SERVER", err)
	}
}
