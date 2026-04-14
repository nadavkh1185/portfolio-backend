package controllers

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File not found",
		})
		return
	}

	// generate filename
	filename := filepath.Base(file.Filename)

	// simpan ke folder uploads
	savePath := "./uploads/" + filename

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to save file",
		})
		return
	}

	// return URL
	c.JSON(http.StatusOK, gin.H{
		"url": "/uploads/" + filename,
	})
}