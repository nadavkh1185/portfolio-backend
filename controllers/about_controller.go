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

	// ambil existing data
	if err := config.DB.First(&about).Error; err != nil {
		about = models.About{}
	}

	// =========================
	// HANDLE TEXT (FORM VALUE)
	// =========================
	subtitle := c.PostForm("subtitle")
	p1 := c.PostForm("paragraph1")
	p2 := c.PostForm("paragraph2")
	p3 := c.PostForm("paragraph3")

	if subtitle != "" {
		about.Subtitle = subtitle
	}
	if p1 != "" {
		about.Paragraph1 = p1
	}
	if p2 != "" {
		about.Paragraph2 = p2
	}
	if p3 != "" {
		about.Paragraph3 = p3
	}

	// =========================
	// HANDLE IMAGE
	// =========================
	file, err := c.FormFile("image")
	if err == nil {
		path := "uploads/" + file.Filename

		if err := c.SaveUploadedFile(file, path); err == nil {
			about.ImageURL = path
		}
	}

	// =========================
	// SAVE DB
	// =========================
	if about.ID == 0 {
		config.DB.Create(&about)
	} else {
		config.DB.Save(&about)
	}

	c.JSON(http.StatusOK, about)
}