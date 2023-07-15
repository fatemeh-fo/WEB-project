package handlers

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

const jwtSecret = "your-secret-key" // Replace this with your own secret key for JWT

// GenerateToken generates a JWT token for the given username.
func GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
