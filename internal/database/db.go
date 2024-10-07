package database

import (
	"log"
	config "my-trips-api/configs"
	"my-trips-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func ConnectionToDatabase() {

	_, err := config.LoadConfig(".")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	DB, err = gorm.Open(postgres.Open(config.BuildStringConnection()))
	if err != nil {
		log.Panic("Error connecting to database")
	}
	DB.AutoMigrate(&models.Testimonial{})
}
