package main

import (
	Controllers "auth-service/controllers"
	"auth-service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	r := gin.Default()
	r.POST("/users", Controllers.CreateUsers)
	r.GET("/users", Controllers.GetUsers)

	r.Run()
}
