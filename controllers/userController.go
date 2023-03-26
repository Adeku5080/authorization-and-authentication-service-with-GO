package controllers

import (
	// "auth-service/helper"
	"auth-service/initializers"
	"auth-service/models"
	// "net/http"

	"github.com/gin-gonic/gin"
)

// func CreateUsers(c *gin.Context) {
// 	//Get data off req body
// 	var body struct {
// 		UserName string
// 		Password string
// 	}
// 	c.Bind(&body)
// 	//create a user
// 	user := models.User{UserName: body.UserName, Password: body.Password}
// 	result := initializers.DB.Create(&user)

// 	if result.Error != nil {
// 		c.Status(400)
// 	}

// 	c.JSON(200, gin.H{
// 		"user": user,
// 	})
// }

func GetUsers(c *gin.Context){
	//all users should be able to perform this task
	var users []models.User
	initializers.DB.Find(&users)

	//return users
		c.JSON(200, gin.H{
		"users": users,
	})
}

func GetUser(c *gin.Context){
	//only admins can perform this action
	// userId := c.Param("id")
	// if err := helper.MatchUserTypeToId(c,userId);err != nil{
	// 	c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
	// 	return
	// }
}

func CreateUser(c *gin.Context){
   //only admin can perform this task
}

func DeleteUser(c *gin.Context){
  //only admin users can perform this
}

