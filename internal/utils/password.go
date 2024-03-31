package utils

import "golang.org/x/crypto/bcrypt"

// GenerateHashPassword generates a hash from the given password using bcrypt.
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// VerifyPassword compares a password with a hashed password.
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
