package user

import (
	"log"
	"net/http"
	"omega/internal/core"
	"strconv"

	"github.com/gin-gonic/gin"
	// "omega/internal/glog"
)

type API struct {
	Service Service
	Engine  core.Engine
}

func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

func (p *API) FindAll(c *gin.Context) {
	users := p.Service.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (p *API) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.Service.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (p *API) Create(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdUser := p.Service.Save(user)

	c.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (p *API) Update(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user = p.Service.FindByID(uint(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Service.Save(user)

	c.Status(http.StatusOK)
}

func (p *API) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.Service.FindByID(uint(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Service.Delete(user)

	c.Status(http.StatusOK)
}
