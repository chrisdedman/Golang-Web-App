package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sandbox-science/deep-focus/config/database"
)

// GenerateToken generates a JWT token for the provided user.
func GenerateToken(user database.User) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	// Set token claims
	claims := jwt.MapClaims{
		"auth": true,
		"id":   user.ID,
		"exp":  time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
		"role": "user",
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
func ValidateToken(c *gin.Context) (jwt.MapClaims, error) {
	token, err := GetToken(c)

	if err != nil {
		return nil, err
	}

	// Check if token is valid
	if user, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return user, nil
	}

	return nil, errors.New("invalid token provided")
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
	// Extract token from request header
	tokenString := getTokenFromRequest(c)

	if tokenString == "" {
		return nil, errors.New("no token provided")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	if token.Valid {
		fmt.Println("Token is valid:", token.Valid)
	} else {
		fmt.Println("Error: Token is invalid:", err)
	}

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("token is malformed")
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				return nil, errors.New("token is expired or not yet valid")
			} else {
				return nil, fmt.Errorf("token validation error: %v", err)
			}
		}
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	return token, nil
}

// getTokenFromRequest extracts the JWT token from the request headers.
func getTokenFromRequest(c *gin.Context) string {
	authHeader := c.Request.Header.Get("Cookie")
	if authHeader == "" {
		return ""
	}

	cookieParts := strings.Split(authHeader, ";")
	var token string
	for _, part := range cookieParts {
		part = strings.TrimSpace(part)
		if strings.HasPrefix(part, "token=") {
			token = strings.TrimPrefix(part, "token=")
			break
		}
	}

	return token
}
