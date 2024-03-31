package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandbox-science/deep-focus/internal/controllers"
	middleware "github.com/sandbox-science/deep-focus/internal/middlewares"
	"gorm.io/gorm"
)

// AuthRoutes registers the authentication routes to the provided Gin router.
func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	server := controllers.NewServer(db)
	route := router.Group("/")
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "In development...",
		})
	})
	route.POST("/login", server.Login)
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	route.POST("/signup", server.Register)
	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
	})
	authorized := route.Group("/api/admin")
	authorized.Use(middleware.JwtAuthMiddleware())

	authorized.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	authorized.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API handler",
		})
	})
	authorized.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// Add a route for the logout endpoint
	route.POST("/logout", server.Logout)
}
