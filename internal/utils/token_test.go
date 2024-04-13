package utils_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-web-app/internal/utils"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	err := godotenv.Load("../../.env.example")
	assert.NoError(t, err)

	token, err := utils.GenerateToken(*user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}

func TestValidateTokenError(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	req.Header.Set("Cookie", "this_is_a_jwt_fake_test_token")

	rr := httptest.NewRecorder()
	r := gin.New()

	r.GET("/user/dashboard", func(c *gin.Context) {
		token, err := utils.ValidateToken(c)
		assert.Error(t, err)
		assert.Nil(t, token)
	})

	r.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnauthorized, rr.Code)
}
