package jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID uint32
	jwt.RegisteredClaims
}

var secretKey = []byte("your-secret-key")

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
