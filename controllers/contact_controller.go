package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"os"
	"portfolio-backend/config"
	"portfolio-backend/models"
	"strings"

	"github.com/gin-gonic/gin"
)

type ContactMessageInput struct {
	Name    string `json:"name" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	Message string `json:"message" binding:"required"`
}

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

func SendContactMessage(c *gin.Context) {
	var input ContactMessageInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Name = strings.TrimSpace(input.Name)
	input.Subject = strings.TrimSpace(input.Subject)
	input.Message = strings.TrimSpace(input.Message)

	if input.Name == "" || input.Subject == "" || input.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name, subject, and message are required"})
		return
	}

	smtpHost := strings.TrimSpace(os.Getenv("SMTP_HOST"))
	smtpPort := strings.TrimSpace(os.Getenv("SMTP_PORT"))
	smtpUsername := strings.TrimSpace(os.Getenv("SMTP_USERNAME"))
	smtpPassword := strings.TrimSpace(os.Getenv("SMTP_PASSWORD"))
	smtpFrom := strings.TrimSpace(os.Getenv("SMTP_FROM"))
	recipient := strings.TrimSpace(os.Getenv("CONTACT_RECEIVER_EMAIL"))

	if recipient == "" {
		recipient = "vinawandakhodijah@gmail.com"
	}

	if smtpHost == "" || smtpPort == "" || smtpUsername == "" || smtpPassword == "" || smtpFrom == "" {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"error": "SMTP configuration is incomplete",
		})
		return
	}

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpHost)

	body := strings.Join([]string{
		fmt.Sprintf("To: %s", recipient),
		fmt.Sprintf("Subject: %s", input.Subject),
		"MIME-Version: 1.0",
		"Content-Type: text/plain; charset=\"UTF-8\"",
		"",
		fmt.Sprintf("Name: %s", input.Name),
		"",
		input.Message,
	}, "\r\n")

	if err := smtp.SendMail(
		fmt.Sprintf("%s:%s", smtpHost, smtpPort),
		auth,
		smtpFrom,
		[]string{recipient},
		[]byte(body),
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send email",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "email sent successfully",
	})
}
