package types

import (
	"github.com/golang-jwt/jwt/v5"
)

// JwtCustomClaim represents the custom JWT claims
type JwtCustomClaim struct {
	UID   int    `json:"uid"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
