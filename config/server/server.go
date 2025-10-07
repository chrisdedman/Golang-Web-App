package server

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/internal/routes"
	"gorm.io/gorm"
)

func Server(host string, db *gorm.DB) {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	// Serve static assets (CSS, JS, images, etc.)
	router.Static("/assets", "./assets")

	// Register all routes
	routes.AuthRoutes(router, db)

	// Load all templates
	files := []string{
		"views/user/authentication/login.html", "views/user/authentication/signup.html",
		"views/user/authentication/logout.html",
		"views/user/account/update.html", "views/user/account/delete.html",
		"views/layout/footer.html", "views/layout/header.html",
		"views/app/dashboard.html", "views/app/index.html", "views/errors.html",
	}

	router.LoadHTMLFiles(files...)

	// Start the server
	router.Run(host)
}
