package controllers

import (
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetSkills(c *gin.Context) {
	var skills []models.Skill
	config.DB.Find(&skills)
	c.JSON(http.StatusOK, skills)
}

func CreateSkill(c *gin.Context) {
	var skill models.Skill

	if err := c.ShouldBindJSON(&skill); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&skill)
	c.JSON(http.StatusOK, skill)
}

func UpdateSkill(c *gin.Context) {
	id := c.Param("id")

	var skill models.Skill
	if err := config.DB.First(&skill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	var input models.Skill
	c.ShouldBindJSON(&input)

	config.DB.Model(&skill).Updates(input)
	c.JSON(http.StatusOK, skill)
}

func DeleteSkill(c *gin.Context) {
	id := c.Param("id")

	var skill models.Skill
	if err := config.DB.First(&skill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	config.DB.Delete(&skill)
	c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}