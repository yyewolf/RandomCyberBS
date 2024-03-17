package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type AuthClaims struct {
	jwt.RegisteredClaims

	Roles []string `json:"roles"`
}
