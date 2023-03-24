package Controllers

import (
	"auth-service/initializers"
	"auth-service/models"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	//Get data off req body
	var body struct {
		UserName string
		Password string
	}
	c.Bind(&body)
	//create a user
	user := models.User{UserName: body.UserName, Password: body.Password}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func GetUsers(c *gin.Context){
	//get users
	var users []models.User
	initializers.DB.Find(&users)

	//return users
		c.JSON(200, gin.H{
		"users": users,
	})
}
