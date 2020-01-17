package models

import "github.com/dgrijalva/jwt-go"

// JWTClaims for JWT
type JWTClaims struct {
	Username string `json:"username"`
	ID       uint64 `json:"id"`
	jwt.StandardClaims
}
