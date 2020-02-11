package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	// "net/http"
	"omega/engine"
	"omega/internal/response"
)

// API for injecting auth service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI for auth used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// Logout delete session // TODO: if not implemented then delete the mehtod
func (p *API) Logout(c *gin.Context) {
	// res := response.Response{Context:c}

	// var auth Auth
	// err := c.ShouldBindJSON(&auth)
	// if err != nil {
	// 	res.Failed(http.StatusUnauthorized, 1401, err.Error(), "")
	// 	return
	// }

	// err = p.Service.Logout(auth)
	// if err != nil {
	// 	res.Failed(http.StatusUnauthorized, 1401, err.Error(), "")
	// 	return
	// }

	// res.Success(http.StatusOK, "logout success", "", 0)
}

// Login user
func (p *API) Login(c *gin.Context) {
	var auth Auth

	if err := c.BindJSON(&auth); err != nil {
		response.ErrorInBinding(c, err, "Login")
		return
	}

	user, err := p.Service.Login(auth)
	if err != nil {
		p.Engine.Record(c, "auth-login-failed", auth.Username, len(auth.Password))
		c.JSON(http.StatusUnauthorized, &response.Result{
			Message: "Username or Password is wrong",
			Code:    1401,
			Error:   err.Error(),
		})
		return
	}

	user.Password = ""
	// user.Extra = ""
	p.Engine.Record(c, "auth-login-success", user)
	response.Success(c, user)
}
