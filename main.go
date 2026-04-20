package main

import (
	"log"
	"os"
	"time"

	"portfolio-backend/config"
	"portfolio-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found (using system env)")
	}

	// Connect DB
	config.ConnectDB()

	// Init Gin
	r := gin.Default()

	// CORS FIX (IMPORTANT)
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000", // frontend kamu
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Origin", "Content-Type", "Authorization",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Static file
	r.Static("/uploads", "./uploads")

	// Routes
	routes.SetupRoutes(r)

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "API running",
		})
	})

	// Port (Railway wajib pakai ENV)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Running on port:", port)
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Failed to run server:", err)
	}
}