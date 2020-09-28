package basapi

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/param"
	"omega/internal/response"

	"github.com/gin-gonic/gin"
)

// AuthAPI for injecting auth service
type AuthAPI struct {
	Service service.BasAuthServ
	Engine  *core.Engine
}

// ProvideAuthAPI for auth used in wire
func ProvideAuthAPI(p service.BasAuthServ) AuthAPI {
	return AuthAPI{Service: p, Engine: p.Engine}
}

// Login auth
func (p *AuthAPI) Login(c *gin.Context) {
	var auth basmodel.Auth
	resp, params := response.NewParam(p.Engine, c, basterm.Users, base.Domain)

	if err := resp.Bind(&auth, "E1053877", base.Domain, basterm.UsernameAndPassword); err != nil {
		return
	}

	user, err := p.Service.Login(auth, params)
	if err != nil {
		resp.Error(err).JSON()
		resp.Record(base.LoginFailed, auth.Username, len(auth.Password))
		return
	}

	tmpUser := user
	tmpUser.Extra = nil

	resp.Record(base.BasLogin, tmpUser)
	resp.Status(http.StatusOK).
		Message(basterm.UserLogedInSuccessfully).
		JSON(user)
}

// Logout will erase the resources from access.Cache
func (p *AuthAPI) Logout(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	params := param.Get(c, p.Engine, basterm.Users)
	p.Service.Logout(params)
	resp.Record(base.BasLogout)
	resp.Status(http.StatusOK).
		Message("user logged out").
		JSON()
}

// TemporaryToken is used for creating temporary access token for download excel and etc
func (p *AuthAPI) TemporaryToken(c *gin.Context) {
	// var auth basmodel.Auth
	resp := response.New(p.Engine, c, base.Domain)

	params := param.Get(c, p.Engine, "users")

	tmpKey, err := p.Service.TemporaryToken(params)
	if err != nil {
		resp.Status(http.StatusInternalServerError).Error(corerr.YouDontHavePermission).JSON()
		return
	}

	resp.Status(http.StatusOK).
		Message(corterm.TemporaryToken).
		JSON(tmpKey)

}
