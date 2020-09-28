package basmid

import (
	"omega/domain/base"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/pkg/limberr"
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
				err := limberr.New("token is required", "E1088822").
					Custom(corerr.UnauthorizedErr).
					Message(corerr.TokenIsRequired).Build()
				response.New(engine, c, base.Domain).Error(err).Abort().JSON()
				return
			}

			token = tokenArr[0][7:]
		}

		claims := &types.JWTClaims{}

		if tkn, err := jwt.ParseWithClaims(token, claims, fJWT); err != nil {
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
			err = limberr.Take(err).Custom(corerr.UnauthorizedErr).
				Message(corerr.TokenIsNotValid).Build()
			response.New(engine, c, base.Domain).Error(err).Abort().JSON()
			return
		}

		err = limberr.Take(err).Custom(corerr.UnauthorizedErr).
			Message(corerr.TokenIsExpired).Build()
		response.New(engine, c, base.Domain).Error(err).Abort().JSON()
		return
	}
}

func checkToken(c *gin.Context, token *jwt.Token, engine *core.Engine) {
	if !token.Valid {
		err := limberr.New(corerr.TokenIsNotValid, "E1054321").Custom(corerr.UnauthorizedErr).
			Message(corerr.TokenIsNotValid).Build()
		response.New(engine, c, base.Domain).Error(err).Abort().JSON()
		return
	}
}
