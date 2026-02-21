package config

import (
	"fmt"
	"log"
	"os"

	"laundry-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	database, err:= gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	//auto migrate
	err = database.AutoMigrate(&models.User{}, &models.LaundryPackage{}, &models.Contact{}, &models.WebContent{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database

	log.Println("Connected to database")
}