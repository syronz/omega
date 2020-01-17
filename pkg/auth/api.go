package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omega/engine"
	"omega/internal/response"
)

type API struct {
	Service Service
	Engine  engine.Engine
}

func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

func (p *API) Logout(c *gin.Context) {
	res := response.Response{Context:c}

	var auth Auth
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		res.Failed(http.StatusUnauthorized, 1401, err.Error(), "")
		return
	}

	err = p.Service.Logout(auth)
	if err != nil {
		res.Failed(http.StatusUnauthorized, 1401, err.Error(), "")
		return
	}

	res.Success(http.StatusOK, "logout success", "", 0)
}

func (p *API) Login(c *gin.Context) {
	res := response.Response{Context:c}

	var auth Auth
	err := c.ShouldBindJSON(&auth)
	if err != nil {
		res.Failed(http.StatusUnauthorized, 1401, err.Error(), "")
		return
	}

	user, err := p.Service.Login(auth)
	if err != nil {
		res.Failed(http.StatusUnauthorized, 1401, err.Error(), "")
		return
	}

	res.Success(http.StatusOK, "login success", user, 0)
}
