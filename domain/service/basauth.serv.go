package service

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/message/baserr"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/limberr"
	"omega/pkg/password"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// BasAuthServ defining auth service
type BasAuthServ struct {
	Engine *core.Engine
}

// ProvideBasAuthService for auth is used in wire
func ProvideBasAuthService(engine *core.Engine) BasAuthServ {
	return BasAuthServ{Engine: engine}
}

// Login User
func (p *BasAuthServ) Login(auth basmodel.Auth, params param.Param) (user basmodel.User, err error) {
	if err = auth.Validate(coract.Login); err != nil {
		err = limberr.Take(err, "E1053212").
			Custom(corerr.ValidationFailedErr).Build()
		return
	}

	jwtKey := p.Engine.Envs.ToByte(base.JWTSecretKey)

	userServ := ProvideBasUserService(basrepo.ProvideUserRepo(p.Engine))
	if user, err = userServ.FindByUsername(auth.Username); err != nil {
		err = limberr.Take(err).Custom(corerr.UnauthorizedErr).
			Message(baserr.UsernameOrPasswordIsWrong).Build()
		return
	}

	if password.Verify(auth.Password, user.Password,
		p.Engine.Envs[base.PasswordSalt]) {

		expirationTime := time.Now().
			Add(p.Engine.Envs.ToDuration(base.JWTExpiration) * time.Second)
		claims := &types.JWTClaims{
			Username: auth.Username,
			ID:       user.ID,
			Lang:     user.Lang,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		var extra struct {
			Token string `json:"token"`
		}
		if extra.Token, err = token.SignedString(jwtKey); err != nil {
			err = limberr.Take(err).Message(corerr.InternalServerError).Build()
			err = corerr.TickCustom(err, corerr.InternalServerErr, "E1042238",
				"error in generating token")
			return
		}

		user.Extra = extra
		user.Password = ""
		BasAccessDeleteFromCache(user.ID)

	} else {
		err = limberr.New("wrong password").Message(baserr.UsernameOrPasswordIsWrong).Build()
		err = corerr.TickCustom(err, corerr.UnauthorizedErr, "E1043108", "wrong password")
	}

	return
}

// Logout erase resources from the cache
func (p *BasAuthServ) Logout(params param.Param) {
	BasAccessResetCache(params.UserID)
}

// TemporaryToken generate instant token for downloading excels and etc
func (p *BasAuthServ) TemporaryToken(params param.Param) (tmpKey string, err error) {
	jwtKey := p.Engine.Envs.ToByte(base.JWTSecretKey)

	expirationTime := time.Now().Add(consts.TemporaryTokenDuration * time.Second)
	claims := &types.JWTClaims{
		ID:   params.UserID,
		Lang: params.Lang,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	if tmpKey, err = token.SignedString(jwtKey); err != nil {
		err = corerr.Tick(err, "E1044682", "temporary token not generated")
		return
	}

	return
}
