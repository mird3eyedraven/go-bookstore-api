package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tkn string) (jwt.MapClaims, error) {
	parsed, err := jwt.Parse(tkn, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		return claims, nil
	}
	return nil, err
}
