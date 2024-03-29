package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getEnvVar(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	return os.Getenv(key)
}

func main() {
	// main is the entry point of the server application.
	// It initializes the server, sets up the routes, and starts listening for incoming requests.
	// If the HOST_ADDR environment variable is not set, it defaults to ":8080".
	listenAddr := getEnvVar("HOST_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}

	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Static("assets", "./assets")
	router.LoadHTMLGlob("templates/*.html")

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
	router.Run(listenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
