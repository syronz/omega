package helper

import (
	"omega/internal/types"

	"github.com/dgrijalva/jwt-go"
)

// ParseJWT decrypt the access_key
func ParseJWT(token string, JWTSecretKey string) (*types.JWTClaims, error) {
	claims := &types.JWTClaims{}
	jwtKey := []byte(JWTSecretKey)
	fJWT := func(token *jwt.Token) (interface{}, error) { return jwtKey, nil }

	var tkn *jwt.Token
	var err error
	if tkn, err = jwt.ParseWithClaims(token, claims, fJWT); err != nil {
		return nil, err
	} else if !tkn.Valid {
		return nil, err
	}

	return claims, err
}
