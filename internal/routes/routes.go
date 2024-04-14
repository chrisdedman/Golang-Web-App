package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang-web-app/internal/controllers"
	middleware "github.com/golang-web-app/internal/middlewares"
	"gorm.io/gorm"
)

func AuthRoutes(router *gin.Engine, db *gorm.DB) {
	server := controllers.NewServer(db)

	router.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "errors.html", gin.H{})
	})

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

		route.GET("/healthcheck", func(c *gin.Context) {
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
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.GET("/dashboard", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(jwt.MapClaims)
			ctx.HTML(http.StatusOK, "dashboard.html", gin.H{
				"user": user["username"],
			})
		})

		authorized.GET("/api", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "API handler",
			})
		})

		authorized.PUT("/update/:user_id", server.UpdateAccount)
		authorized.GET("/update", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(jwt.MapClaims)
			ctx.HTML(http.StatusOK, "update.html", gin.H{
				"userID": user["id"].(float64),
				"user":   user["username"],
			})
		})

		authorized.DELETE("/delete/:user_id", server.DeleteUser)
		authorized.GET("/delete", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(jwt.MapClaims)
			ctx.HTML(http.StatusOK, "delete.html", gin.H{
				"userID": user["id"].(float64),
				"user":   user["username"],
			})
		})

		authorized.POST("/logout", server.Logout)
		authorized.GET("/logout", func(ctx *gin.Context) {
			user := ctx.MustGet("user").(jwt.MapClaims)
			ctx.HTML(http.StatusOK, "logout.html", gin.H{
				"user": user["username"],
			})
		})
	}
}
