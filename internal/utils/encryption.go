package utils

import (
	"github.com/golang-web-app/internal/models"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the user's password using bcrypt.
func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// VerifyPassword compares a password with a hashed password.
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
