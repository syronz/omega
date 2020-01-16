package auth

import (
	"net/http"
	"omega/internal/core"

	"github.com/gin-gonic/gin"
)

type API struct {
	Service Service
	Engine  core.Engine
}

func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

func (p *API) Logout(c *gin.Context) {
	users := p.Service.Logout()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (p *API) Login(c *gin.Context) {
	var auth Auth
	err := c.BindJSON(&auth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	auth, _ = p.Service.Login(auth)

	c.JSON(http.StatusOK, gin.H{"auth": auth})
}
