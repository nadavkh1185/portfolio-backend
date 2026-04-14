package controllers

import (
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAbout(c *gin.Context) {
	var about models.About

	if err := config.DB.First(&about).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "About not found"})
		return
	}

	c.JSON(http.StatusOK, about)
}

func UpdateAbout(c *gin.Context) {
	var about models.About
	var input models.About

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&about).Error; err != nil {
		config.DB.Create(&input)
		c.JSON(http.StatusOK, input)
		return
	}

	config.DB.Model(&about).Updates(input)

	c.JSON(http.StatusOK, about)
}