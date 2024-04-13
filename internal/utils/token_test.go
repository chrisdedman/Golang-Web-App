package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/internal/models"
	"github.com/golang-web-app/internal/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	err := godotenv.Load("../../.env.example")
	assert.NoError(t, err)

	token, err := utils.GenerateToken(models.User{
		Username: "testDev",
		Email:    "testdev@gmail.com",
		Password: "Password",
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateToken(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	// Set the token in the request header
	req.Header.Set("Cookie", "this_is_a_jwt_fake_test_token")

	// Create a ResponseRecorder (which implements http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	r := gin.New()

	// Define a route that calls ValidateToken
	r.GET("/", func(c *gin.Context) {
		user, err := utils.ValidateToken(c)
		assert.Error(t, err)
		assert.Nil(t, user)
	})

	// Perform the request
	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
}
