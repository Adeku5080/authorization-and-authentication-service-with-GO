package controllers

import (
	"auth-service/initializers"
	"auth-service/models"
	"net/http"
	"os"
	"time"

	// "auth-service/helpers"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

var validate = validator.New()

func HashPassword() {}

func VerifyPassword() {}

func RegisterUser(c *gin.Context) {

	var body struct {
		UserName string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}

	//hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})
		return
	}

	user := models.User{UserName: body.UserName, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
	}

	c.JSON(200, gin.H{
		"user": user,
	})

}

func LoginUser(c *gin.Context) {
	var body struct {
		UserName string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to read body",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "UserName = ?", body.UserName)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid username or passwors",
		})
		return
	}

	//compare sent in password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"eror": "Invalid email or password",
		})
		return
	}

	//generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": "user.ID",
		"exp": time.Now().Add(time.Hour).Unix()
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(os.Getenv("SECRET")) 

	if( err != nil){
		c.JSON(http.statusBadRequest,gin.H{
			"error" : "Invalid credentials"
		})
	}

	//send the token back
	c.JSON(http.statusOK,gin.H{
		"token" : tokenString
	})
}
