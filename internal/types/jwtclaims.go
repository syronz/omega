package types

import (
	"omega/pkg/dict"

	"github.com/dgrijalva/jwt-go"
)

// JWTClaims for JWT
type JWTClaims struct {
	Username string        `json:"username"`
	ID       RowID         `json:"id"`
	Language dict.Language `json:"language"`
	jwt.StandardClaims
}
