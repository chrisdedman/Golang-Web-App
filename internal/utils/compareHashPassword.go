package utils

import "golang.org/x/crypto/bcrypt"

// CompareHashPassword compares a plain-text password with a hashed password and returns true if they match.
// It uses bcrypt.CompareHashAndPassword to perform the comparison.
func CompareHashPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
