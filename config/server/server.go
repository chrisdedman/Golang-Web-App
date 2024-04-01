package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sandbox-science/deep-focus/config/database"
	"github.com/sandbox-science/deep-focus/internal/routes"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Load the environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database configuration
	config := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Set the listen address
	listenAddr := os.Getenv("HOST_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	// Initialize the client IP address
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	// Serve static files and templates
	router.Static("assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

	// Initialize the database
	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Initialize the routes and run the server
	routes.AuthRoutes(router, db)
	router.Run(listenAddr)
}
