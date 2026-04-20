package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"portfolio-backend/config"
	"portfolio-backend/models"

	"github.com/gin-gonic/gin"
)

// ======================
// GET ALL
// ======================
func GetSkills(c *gin.Context) {
	var skills []models.Skill
	config.DB.Find(&skills)
	c.JSON(http.StatusOK, skills)
}


// ======================
// GET BY ID (WAJIB TAMBAH)
// ======================
func GetSkillByID(c *gin.Context) {
	id := c.Param("id")

	var skill models.Skill
	if err := config.DB.First(&skill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	c.JSON(http.StatusOK, skill)
}


// ======================
// CREATE (MULTIPART)
// ======================
func CreateSkill(c *gin.Context) {
	var skill models.Skill

	// ambil text
	title := c.PostForm("image_title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "image_title is required"})
		return
	}
	skill.ImageTitle = title

	// ambil file
	file, err := c.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		path := "uploads/" + filename

		if err := c.SaveUploadedFile(file, path); err == nil {
			skill.ImageURL = path
		}
	}

	config.DB.Create(&skill)
	c.JSON(http.StatusOK, skill)
}


// ======================
// UPDATE (MULTIPART)
// ======================
func UpdateSkill(c *gin.Context) {
	id := c.Param("id")

	var skill models.Skill
	if err := config.DB.First(&skill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	// update title
	title := c.PostForm("image_title")
	if title != "" {
		skill.ImageTitle = title
	}

	// update image (optional)
	file, err := c.FormFile("image")
	if err == nil {
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), filepath.Base(file.Filename))
		path := "uploads/" + filename

		if err := c.SaveUploadedFile(file, path); err == nil {
			skill.ImageURL = path
		}
	}

	config.DB.Save(&skill)
	c.JSON(http.StatusOK, skill)
}


// ======================
// DELETE
// ======================
func DeleteSkill(c *gin.Context) {
	id := c.Param("id")

	var skill models.Skill
	if err := config.DB.First(&skill, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Skill not found"})
		return
	}

	config.DB.Delete(&skill)
	c.JSON(http.StatusOK, gin.H{"success": true})
}