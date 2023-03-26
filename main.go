package main

import (
    "auth-service/controllers" 
	"auth-service/initializers"
 
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	r := gin.Default()

	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)

	r.Run()

}
