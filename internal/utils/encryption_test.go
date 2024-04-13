package utils_test

import (
	"testing"

	"github.com/golang-web-app/internal/models"
	"github.com/golang-web-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

var user = &models.User{
	Username: "dev",
	Password: "password123",
	Email:    "dev@example.com",
}

/*
 * Test cases for the HashPassword function
 */
func TestHashPasswordNoError(t *testing.T) {
	err := utils.HashPassword(user)
	assert.NoError(t, err, "No error while hashing password")
}

func TestHashPasswordNil(t *testing.T) {
	err := utils.HashPassword(user)
	assert.Nil(t, err, "Err should be nil")
}

/*
 * Test cases for the VerifyPassword function
 */
func TestVerifyPasswordNoError(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	assert.NoError(t, err, "No error while hashing password")

	err = utils.VerifyPassword(user.Password, string(hashedPassword))
	assert.NoError(t, err, "No error while verifying password")
}

func TestVerifyPasswordError(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password1"), bcrypt.DefaultCost)
	assert.NoError(t, err, "No error while hashing password")

	err = utils.VerifyPassword(user.Password, string(hashedPassword))
	assert.Error(t, err, "Error while verifying password")
}
