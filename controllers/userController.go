package Controllers

import "github.com/gin-gonic/gin"

func createUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
