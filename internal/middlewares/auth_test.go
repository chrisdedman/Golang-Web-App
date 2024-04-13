package middleware_test

import (
	"testing"

	middleware "github.com/golang-web-app/internal/middlewares"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	auth := middleware.AuthMiddleware()
	assert.NotEmpty(t, auth)
}
