package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func apiHandler(w http.ResponseWriter) {
	// apiHandler is an HTTP handler that writes a JSON response to the given ResponseWriter.
	// The response contains a message "API handler".
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"message": "API handler"}`)
}

func main() {
	// main is the entry point of the server application.
	// It initializes the server, sets up the routes, and starts listening for incoming requests.
	// If the HOST_ADDR environment variable is not set, it defaults to ":8080".
	listenAddr := os.Getenv("HOST_ADDR")
	if len(listenAddr) == 0 {
		listenAddr = ":8080"
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/api", func(c *gin.Context) {
		apiHandler(c.Writer)
	})
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the server!",
		})
	})
	r.Run(listenAddr) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
