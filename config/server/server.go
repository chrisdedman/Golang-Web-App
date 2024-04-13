package server

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/config/database"
	"github.com/golang-web-app/internal/routes"
	"github.com/joho/godotenv"
)

func Server() {
	router := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	host := os.Getenv("HOST_ADDR")
	if len(host) == 0 {
		host = ":8080"
	}

	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	router.Static("assets", "./assets")

	files := []string{
		"views/user/authentication/login.html", "views/user/authentication/signup.html",
		"views/user/authentication/logout.html", "views/user/account/update.html",
		"views/user/account/delete.html",
		"views/layout/footer.html", "views/layout/header.html",
		"views/app/dashboard.html", "views/app/index.html", "views/errors.html",
	}
	router.LoadHTMLFiles(files...)

	db, err := database.InitDB(config)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	routes.AuthRoutes(router, db)
	router.Run(host)
}
