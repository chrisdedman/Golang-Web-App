package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandbox-science/deep-focus/internal/controllers"
)

// AuthRoutes registers the authentication routes to the provided Gin router.
func AuthRoutes(router *gin.Engine) {
	router.POST("/login", controllers.Login)
	router.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{})
	})

	router.POST("/signup", controllers.Signup)
	router.GET("/signup", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "signup.html", gin.H{})
	})

	router.POST("/logout", controllers.Logout)
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"content": "In development...",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "API handler",
		})
	})
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
}
