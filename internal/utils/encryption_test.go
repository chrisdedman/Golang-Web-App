package utils_test

import (
	"testing"

	"github.com/golang-web-app/internal/models"
	"github.com/golang-web-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestHashPassword(t *testing.T) {
	user := &models.User{
		Password: "password123",
	}

	err := utils.HashPassword(user)
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password123"))
	assert.NoError(t, err)

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("password1234"))
	assert.Error(t, err)
}

func TestVerifyPassword(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	assert.NoError(t, err)

	err = utils.VerifyPassword("password123", string(hashedPassword))
	assert.NoError(t, err)

	err = utils.VerifyPassword("password1234", string(hashedPassword))
	assert.Error(t, err)
}
