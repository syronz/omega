package user

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserAPI struct {
	UserService UserService
}

func ProvideUserAPI(p UserService) UserAPI {
	return UserAPI{UserService: p}
}

func (p *UserAPI) FindAll(c *gin.Context) {
	users := p.UserService.FindAll()

	c.JSON(http.StatusOK, gin.H{"users": ToUserDTOs(users)})
}

func (p *UserAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"user": ToUserDTO(user)})
}

func (p *UserAPI) Create(c *gin.Context) {
	var userDTO UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdUser := p.UserService.Save(ToUser(userDTO))

	c.JSON(http.StatusOK, gin.H{"user": ToUserDTO(createdUser)})
}

func (p *UserAPI) Update(c *gin.Context) {
	var userDTO UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	user := p.UserService.FindByID(uint(id))
	if user == (User{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	user.Code = userDTO.Code
	user.Price = userDTO.Price
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
