// PATH: deep-focus/utils/ParseToken.go

package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/sandbox-science/deep-focus/internal/database"
)

func ParseToken(tokenString string) (claims *database.Claims, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &database.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*database.Claims)

	if !ok {
		return nil, err
	}

	return claims, nil
}
