package types

import (
	"github.com/dgrijalva/jwt-go"
)

// JWTClaims for JWT
type JWTClaims struct {
	Username string `json:"username"`
	ID       RowID  `json:"id"`
	Language string `json:"language"`
	jwt.StandardClaims
}
