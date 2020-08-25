package basapi

import (
	"net/http"
	"omega/domain/base/basevent"
	"omega/domain/base/basmodel"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/response"
	"omega/internal/term"

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
	resp := response.New(p.Engine, c)

	if err := c.BindJSON(&auth); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	user, err := p.Service.Login(auth)
	if err != nil {
		resp.Error(err).JSON()
		resp.Record(term.Auth_login_failed, auth.Username, len(auth.Password))
		return
	}

	userTmp := user
	userTmp.Extra = nil

	resp.Record(basevent.BasLogin, nil, userTmp)
	resp.Status(http.StatusOK).
		Message(term.User_loged_in_successfully).
		JSON(user)
}

// Logout will erase the resources from access.Cache
func (p *AuthAPI) Logout(c *gin.Context) {
	resp := response.New(p.Engine, c)
	params := param.Get(c, p.Engine, thisUsers)
	p.Engine.Debug(params.UserID)
	p.Service.Logout(params)
	resp.Record(basevent.BasLogout)
	resp.Status(http.StatusOK).
		Message("user logged out").
		JSON()
}

// TemporaryToken is used for creating temporary access token for download excel and etc
func (p *AuthAPI) TemporaryToken(c *gin.Context) {
	// var auth basmodel.Auth
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, "users")

	tmpKey, err := p.Service.TemporaryToken(params)
	if err != nil {
		resp.Status(http.StatusInternalServerError).Error(term.You_dont_have_permission).JSON()
		return
	}

	resp.Status(http.StatusOK).
		Message(term.Temporary_Token).
		JSON(tmpKey)

}
