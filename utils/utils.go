package utils

import (
	"fitness-api/types"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Generates a new token and returns the string of the token
func GenToken(uid int, name, email string) (string, error) {
	secretKey := []byte(os.Getenv("SECRET_KEY"))

	claims := &types.JwtCustomClaim{
		UID:   uid,
		Name:  name,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
