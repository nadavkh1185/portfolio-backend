package config

import (
	"log"
	"os"
	"portfolio-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}

	DB = db

	DB.AutoMigrate(
		&models.Profile{},
		&models.About{},
		&models.Skill{},
		&models.Project{},
		&models.Experience{},
		&models.Contact{},
	)

	log.Println("Database connected & migrated")
}