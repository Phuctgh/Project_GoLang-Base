// cmd/main.go
package main

import (
	"log"

	"auth-service/internal/config"
	"auth-service/internal/handler"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// verify environment variables
	if os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_USER") == "" || os.Getenv("DB_PASSWORD") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("Some DB environment variables are missing")
	}

	// connect to database
	config.SetupDatabase()
	// make sure to close the database connection when the application exits
	defer config.CloseDB()

	// initialize the router
	router := gin.Default()

	// register routes
	handler.RegisterAuthRoutes(router)

	// run the server
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
