package main

import (
	"log"
	"os"

	"github.com/golang-web-app/config/database"
	"github.com/golang-web-app/config/server"
	"github.com/golang-web-app/internal/models"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	host := os.Getenv("HOST_ADDR")
	if len(host) == 0 {
		host = ":8080"
	}

	config := models.DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	server.Server(host, db)
}
