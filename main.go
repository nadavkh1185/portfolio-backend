package main

import (
	"log"
	"os"
	"strings"
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

	allowedOrigins := []string{
		"http://localhost:3000",
		"http://127.0.0.1:3000",
		"https://portfolio-frontend-tan-three.vercel.app",
	}

	if frontendURL := strings.TrimSpace(os.Getenv("FRONTEND_URL")); frontendURL != "" {
		allowedOrigins = append(allowedOrigins, frontendURL)
	}

	if extraOrigins := strings.TrimSpace(os.Getenv("FRONTEND_ORIGINS")); extraOrigins != "" {
		for _, origin := range strings.Split(extraOrigins, ",") {
			origin = strings.TrimSpace(origin)
			if origin != "" {
				allowedOrigins = append(allowedOrigins, origin)
			}
		}
	}

	// CORS FIX (IMPORTANT)
	r.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
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
