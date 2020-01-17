package middleware

import (
	"net/http"
	"omega/engine"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"omega/internal/models"
)

// CheckToken is used for decode the token and get public and private information
func CheckToken(e engine.Engine) gin.HandlerFunc {
	jwtKey := []byte(e.Environments.Setting.JWTSecretKey)

	return func(c *gin.Context) {

		tokenArr, ok := c.Request.Header["Authorization"]
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is required"})
			return
		}

		token := tokenArr[0][7:]

		claims := &models.JWTClaims{}

		tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !tkn.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			return
		}

		c.Set("USERNAME", claims.Username)
		c.Set("USER_ID", claims.ID)

		c.Next()
	}
}
