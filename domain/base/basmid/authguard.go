package basmid

import (
	"net/http"
	"omega/domain/base"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/pkg/glog"
	"strings"

	"omega/internal/response"
	"omega/internal/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// AuthGuard is used for decode the token and get public and private information
func AuthGuard(engine *core.Engine) gin.HandlerFunc {
	jwtKey := []byte(engine.Envs[base.JWTSecretKey])
	fJWT := func(token *jwt.Token) (interface{}, error) { return jwtKey, nil }

	return func(c *gin.Context) {

		token := strings.TrimSpace(c.Query("temporary_token"))
		if token == "" {
			tokenArr, ok := c.Request.Header["Authorization"]
			if !ok || len(tokenArr[0]) == 0 {
				response.New(engine, c).Status(http.StatusUnauthorized).Abort().
					Error(corerr.Token_is_required).JSON()
				return
			}

			token = tokenArr[0][7:]
		}

		claims := &types.JWTClaims{}

		if tkn, err := jwt.ParseWithClaims(token, claims, fJWT); err != nil {
			glog.Debug(err)
			checkErr(c, err, engine)
			return
		} else if !tkn.Valid {
			checkToken(c, tkn, engine)
			return
		}

		c.Set("USERNAME", claims.Username)
		c.Set("USER_ID", claims.ID)
		c.Set("LANGUAGE", claims.Lang)
		c.Next()
	}
}

func checkErr(c *gin.Context, err error, engine *core.Engine) {
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			response.New(engine, c).Status(http.StatusUnauthorized).Abort().
				Error(err).JSON()
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
}

func checkToken(c *gin.Context, token *jwt.Token, engine *core.Engine) {
	if !token.Valid {
		response.New(engine, c).Status(http.StatusUnauthorized).Abort().
			Error(corerr.Token_is_not_valid).JSON()
		return
	}
}
