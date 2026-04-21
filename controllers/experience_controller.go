package controllers

import (
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetExperience(c *gin.Context) {
	var experiences []models.Experience
	config.DB.Find(&experiences)
	c.JSON(http.StatusOK, experiences)
}

func CreateExperience(c *gin.Context) {
	var experience models.Experience

	if err := c.ShouldBindJSON(&experience); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&experience)
	c.JSON(http.StatusOK, experience)
}

func GetExperienceByID(c *gin.Context) {
	id := c.Param("id")

	var experience models.Experience
	if err := config.DB.First(&experience, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Experience not found"})
		return
	}

	c.JSON(http.StatusOK, experience)
}

func UpdateExperience(c *gin.Context) {
	id := c.Param("id")

	var experience models.Experience
	if err := config.DB.First(&experience, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Experience not found"})
		return
	}

	var input models.Experience
	c.ShouldBindJSON(&input)

	config.DB.Model(&experience).Updates(input)
	c.JSON(http.StatusOK, experience)
}

func DeleteExperience(c *gin.Context) {
	id := c.Param("id")

	var experience models.Experience
	if err := config.DB.First(&experience, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Experience not found"})
		return
	}

	config.DB.Delete(&experience)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}