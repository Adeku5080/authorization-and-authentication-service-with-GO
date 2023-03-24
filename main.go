package main

import (
	"auth-service/initializers"

	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	r := gin.Default()
	r.GET("/" 
	r.Run() // listen and serve on 0.0.0.0:8080
}
