package controllers

import (
	"net/http"
	"os"
	"portfolio-backend/config"
	"portfolio-backend/models"
	"portfolio-backend/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	if err := config.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
		return
	}

	// sementara plain text (nanti kita upgrade bcrypt)
	if err := utils.CheckPassword(user.Password, input.Password); err != nil {
	c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid credentials"})
	return
}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	secret := os.Getenv("JWT_SECRET")

	tokenString, _ := token.SignedString([]byte(secret))

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}