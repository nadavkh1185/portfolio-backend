package controllers

import (
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetContact(c *gin.Context) {
	var contact models.Contact

	if err := config.DB.First(&contact).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Contact not found"})
		return
	}

	c.JSON(http.StatusOK, contact)
}

func UpdateContact(c *gin.Context) {
	var contact models.Contact
	var input models.Contact

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.First(&contact).Error; err != nil {
		config.DB.Create(&input)
		c.JSON(http.StatusOK, input)
		return
	}

	config.DB.Model(&contact).Updates(input)

	c.JSON(http.StatusOK, contact)
}