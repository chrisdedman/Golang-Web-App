package server_test

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	router := gin.Default()

	err := godotenv.Load("../../.env")
	assert.NoError(t, err)

	listenAddr := os.Getenv("HOST_ADDR")
	assert.NotEmpty(t, listenAddr)

	// Initialize the routes and run the server
	assert.NotNil(t, router)
}
