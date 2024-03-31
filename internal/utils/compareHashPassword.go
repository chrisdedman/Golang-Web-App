package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashPassword compares a plain-text password with a hashed password and returns true if they match.
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
