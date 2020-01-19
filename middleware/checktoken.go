package middleware

import (
	"net/http"
	"omega/engine"

	"omega/internal/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// CheckToken is used for decode the token and get public and private information
func CheckToken(e engine.Engine) gin.HandlerFunc {
	jwtKey := []byte(e.Environments.Setting.JWTSecretKey)
	fJWT := func(token *jwt.Token) (interface{}, error) { return jwtKey, nil }

	return func(c *gin.Context) {
		tokenArr, ok := c.Request.Header["Authorization"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			return
		}
		token := tokenArr[0][7:]
		claims := &models.JWTClaims{}

		if tkn, err := jwt.ParseWithClaims(token, claims, fJWT); err != nil {
			checkErr(c, err)
			return
		} else if !tkn.Valid {
			checkToken(c, tkn)
			return
		}
		c.Set("USERNAME", claims.Username)
		c.Set("USER_ID", claims.ID)
		c.Next()
	}
}

func checkErr(c *gin.Context, err error) {
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}

func checkToken(c *gin.Context, token *jwt.Token) {
	if !token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
		return
	}
}
