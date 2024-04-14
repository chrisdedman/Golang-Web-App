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
	"github.com/golang-web-app/internal/models"
)

func GenerateToken(user models.User) (string, error) {
	tokenLifespan, err := strconv.Atoi(os.Getenv("TOKEN_HOUR_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"auth":     true,
		"id":       user.ID,
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix(),
		"role":     "user",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("API_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(c *gin.Context) (jwt.MapClaims, error) {
	token, err := GetToken(c)

	if err != nil {
		return nil, err
	}

	if user, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return user, nil
	}

	return nil, errors.New("invalid token provided")
}

func GetToken(c *gin.Context) (*jwt.Token, error) {
	tokenString, err := GetTokenFromRequest(c)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("API_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func GetTokenFromRequest(c *gin.Context) (string, error) {
	authHeader := c.Request.Header.Get("Cookie")
	if authHeader == "" {
		return "", errors.New("no cookie provided")
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

	return token, nil
}
