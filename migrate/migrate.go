package main

import (
	"auth-service/initializers"
	"auth-service/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.Connect()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
 