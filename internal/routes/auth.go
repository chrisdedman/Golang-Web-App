package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sandbox-science/deep-focus/internal/controllers"
	middleware "github.com/sandbox-science/deep-focus/internal/middlewares"
	"gorm.io/gorm"
)

// AuthRoutes registers the authentication routes to the provided Gin router.
func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	// Create a new server instance
	server := controllers.NewServer(db)

	/*
		Register the authentication routes. Not protected by JWT middleware.
		The following routes are accesible by anyone.
	*/
	route := router.Group("/")
	{
		route.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{})
		})

		route.POST("/login", server.Login)
		route.GET("/login", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "login.html", gin.H{})
		})

		route.POST("/signup", server.Register)
		route.GET("/signup", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "signup.html", gin.H{})
		})

		route.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		})
	}

	/*
		Authorized routes. Protected by JWT middleware.
		The following routes are only accesible by authenticated users.
	*/
	authorized := route.Group(("/user"))
	authorized.Use(middleware.JwtAuthMiddleware())
	{
		// TO-DO: Add user ID to the context for user deletion
		authorized.GET("/dashboard", func(ctx *gin.Context) {
			userMap := ctx.MustGet("user").(jwt.MapClaims)
			userID := userMap["id"].(float64)

			ctx.HTML(http.StatusOK, "dashboard.html", gin.H{
				"userID": userID,
			})
		})

		authorized.GET("/api", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API handler",
			})
		})

		authorized.DELETE("/delete/:user_id", server.DeleteUser)

		// Protected logout route
		authorized.POST("/logout", server.Logout)
		authorized.GET("/logout", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "logout.html", gin.H{})
		})

		// Wildcard route for default HTML layout
		router.NoRoute(func(c *gin.Context) {
			c.HTML(http.StatusNotFound, "noFound.html", gin.H{})
		})
	}
}
