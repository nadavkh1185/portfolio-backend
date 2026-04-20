package controllers

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

func GetProjects(c *gin.Context) {
	var projects []models.Project

	if err := config.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch projects",
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func saveImage(c *gin.Context, fileHeader *multipart.FileHeader) (string, error) {
	filename := filepath.Base(fileHeader.Filename)
	path := "uploads/" + filename

	if err := c.SaveUploadedFile(fileHeader, path); err != nil {
		return "", err
	}

	return path, nil
}

func GetProjectByID(c *gin.Context) {
	id := c.Param("id")

	var project models.Project
	if err := config.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Project not found",
		})
		return
	}

	c.JSON(http.StatusOK, project)
}

func CreateProject(c *gin.Context) {
	subtitle := c.PostForm("subtitle")

	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Image is required",
		})
		return
	}

	imagePath, err := saveImage(c, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to upload image",
		})
		return
	}

	project := models.Project{
		Subtitle: subtitle,
		ImageURL: imagePath,
	}

	if err := config.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create project",
		})
		return
	}

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

	subtitle := c.PostForm("subtitle")

	// update subtitle kalau ada
	if subtitle != "" {
		project.Subtitle = subtitle
	}

	// cek apakah ada file baru
	file, err := c.FormFile("image")
	if err == nil && file != nil {
		imagePath, err := saveImage(c, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Failed to upload image",
			})
			return
		}
		project.ImageURL = imagePath
	}

	if err := config.DB.Save(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update project",
		})
		return
	}

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

	if err := config.DB.Delete(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete project",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deleted successfully",
	})
}