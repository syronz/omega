package service

import (
	"errors"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/param"
	"omega/internal/term"
	"omega/internal/types"
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
	if err = auth.Validate(coract.Login, params); err != nil {
		return
	}

	jwtKey := p.Engine.Envs.ToByte(base.JWTSecretKey)

	userServ := ProvideBasUserService(basrepo.ProvideUserRepo(p.Engine))
	if user, err = userServ.FindByUsername(auth.Username); err != nil {
		err = errors.New(term.Username_or_password_is_wrong)
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
			return
		}

		user.Extra = extra
		user.Password = ""
		BasAccessDeleteFromCache(user.ID)

	} else {
		err = errors.New(term.Username_or_password_is_wrong)
	}

	return
}

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
		return
	}

	return
}
