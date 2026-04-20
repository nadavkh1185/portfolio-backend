package main

import (
	"log"
	"os"
	"portfolio-backend/config"
	"portfolio-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env (optional)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (using system env)")
	}

	// Connect DB
	config.ConnectDB()

	// Init Gin
	r := gin.Default()

r.Use(cors.New(cors.Config{
	AllowOrigins:     []string{"http://localhost:3000"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
}))

	r.Static("/uploads", "./uploads")

	routes.SetupRoutes(r)

	// Test endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API running",
		})
	})

	// PORT FIX (IMPORTANT)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Running on port:", port)
	r.Run(":" + port)
}
