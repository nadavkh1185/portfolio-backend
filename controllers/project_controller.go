package controllers

import (
	"net/http"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project

	config.DB.Find(&projects)

	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {
	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	config.DB.Create(&project)

	c.JSON(http.StatusOK, project)
}

func UpdateProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project

	if err := config.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Project not found",
		})
		return
	}

	var input models.Project
	c.ShouldBindJSON(&input)

	config.DB.Model(&project).Updates(input)

	c.JSON(http.StatusOK, project)
}

func DeleteProject(c *gin.Context) {
	id := c.Param("id")

	var project models.Project

	if err := config.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Project not found",
		})
		return
	}

	config.DB.Delete(&project)

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}