package controllers

import (
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	var profile models.Profile

	result := config.DB.First(&profile)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Profile not found",
		})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func UpdateProfile(c *gin.Context) {
	var profile models.Profile
	var input models.Profile

	// bind JSON dari request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// cek apakah profile sudah ada
	result := config.DB.First(&profile)

	if result.Error != nil {
		// kalau belum ada → create
		config.DB.Create(&input)
		c.JSON(http.StatusOK, input)
		return
	}

	// kalau sudah ada → update
	config.DB.Model(&profile).Updates(input)

	c.JSON(http.StatusOK, profile)
}