package config

import (
	"fmt"
	"log"
	"os"

	"auth-service/internal/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Global variable for DB connection
var DB *gorm.DB

// SetupDatabase is used to initialize the database connection
func SetupDatabase() {
	var err error
	// Setup the connection string to PostgreSQL
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	// Open connection to database
	DB, err = gorm.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Automigrate: automatically create or update if the table structure changes
	DB.AutoMigrate(&model.User{})
}

// CloseDB will close the DB connection when the application exits
func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Fatalf("Failed to close the database connection: %v", err)
	}
}
