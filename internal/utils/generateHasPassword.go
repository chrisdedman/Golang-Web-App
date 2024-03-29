package utils

import "golang.org/x/crypto/bcrypt"

// GenerateHashPassword generates a hash from the given password using bcrypt.
// It takes a password string as input and returns the generated hash as a string.
// If an error occurs during the hash generation process, it is also returned.
func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
