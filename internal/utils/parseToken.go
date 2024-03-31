package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4" // Import v4 of the JWT library
	"github.com/sandbox-science/deep-focus/internal/database"
)

// GenerateToken generates a JWT token for the provided user.
func GenerateToken(user database.User) (string, error) {
	// Parse token lifespan from environment variable
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	// Create token claims
	claims := jwt.MapClaims{
		"authorized": true,
		"id":         user.ID,
		"exp":        time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
	}

	// Generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken validates the JWT token provided in the request headers.
func ValidateToken(c *gin.Context) error {
	// Extract token from request
	token, err := GetToken(c)
	if err != nil {
		return err
	}

	fmt.Println("Retrieved token:", token)

	// Check if token is valid
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

// GetToken extracts the JWT token from the request headers.
func GetToken(c *gin.Context) (*jwt.Token, error) {
	// Extract token from request headers
	tokenString := getTokenFromRequest(c)
	fmt.Println("Retrieved token:", c.Request.Header)

	// Check if token string is empty
	if tokenString == "" {
		return nil, errors.New("no token provided")
	}

	// Parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Return secret key for validation
		return []byte(os.Getenv("API_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

// getTokenFromRequest extracts the JWT token from the request headers.
func getTokenFromRequest(c *gin.Context) string {
	// Extract token from "Authorization" header
	fmt.Println("--", c.Request.Header)

	// Print raw request
	rawRequest, _ := c.GetRawData()
	fmt.Println("ROW", string(rawRequest))
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}

	// Check if header starts with "Bearer "
	parts := strings.Split(authHeader, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}

	return ""
}
