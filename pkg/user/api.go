package user

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// "omega/internal/glog"
)

type UserAPI struct {
	UserService UserService
}

func ProvideUserAPI(p UserService) UserAPI {
	return UserAPI{UserService: p}
}

func (p *UserAPI) FindAll(c *gin.Context) {
	users := p.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": users})
}

func (p *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (p *UserAPI) Create(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdUser := p.UserService.Save(user)
	// glog.Debug(createdUser)

	c.JSON(http.StatusOK, gin.H{"user": createdUser})
}

func (p *UserAPI) Update(c *gin.Context) {
	var user User
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user = p.UserService.FindByID(uint(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.UserService.Save(user)

	c.Status(http.StatusOK)
}

func (p *UserAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.UserService.Delete(user)

	c.Status(http.StatusOK)
}
