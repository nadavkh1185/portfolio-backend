package main

import (
	"log"
	"portfolio-backend/config"
	"portfolio-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect DB
	config.ConnectDB()

	// Init Gin
	r := gin.Default()

	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	// Test endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API running",
		})
	})

	// Run server
	r.Run(":8080")
}