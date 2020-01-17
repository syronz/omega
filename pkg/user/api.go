package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omega/engine"
	"omega/internal/response"
	"strconv"
)

// API for injecting user service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI for user is used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// FindAll users
func (p *API) FindAll(c *gin.Context) {
	users, err := p.Service.FindAll()

	if err != nil {
		response.RecordNotFound(c, err, "users")
		return
	}

	response.Success(c, users)
}

// FindByID is used for fetch a user by his id
func (p *API) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	user, err := p.Service.FindByID(id)

	if err != nil {
		response.RecordNotFound(c, err, "user")
		return
	}

	response.Success(c, user)
}

// Create user
func (p *API) Create(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)
	if err != nil {
		response.ErrorInBinding(c, err, "create user")
		return
	}

	createdUser, err := p.Service.Save(user)
	if err != nil {
		response.ErrorOnSave(c, err, "user")
		return
	}

	response.Success(c, createdUser)
}

// Update user
func (p *API) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var user User

	err = c.BindJSON(&user)
	if err != nil {
		response.ErrorInBinding(c, err, "update user")
		return
	}
	user.ID = id

	updatedUser, err := p.Service.Save(user)
	if err != nil {
		response.ErrorOnSave(c, err, "update user")
		return
	}

	response.Success(c, updatedUser)
}

// Delete user
func (p *API) Delete(c *gin.Context) {
	// res := response.Response{Context: c}
	// id, _ := strconv.Atoi(c.Param("id"))
	// user, err := p.Service.FindByID(uint(id))
	// if err != nil {
	// 	res.Failed(http.StatusBadRequest, 1400, "user does not exist", "")
	// 	return
	// }

	// err = p.Service.Delete(user)
	// if err != nil {
	// 	res.Failed(http.StatusBadRequest, 1400, "something went wrong, cannot delete this user", "")
	// 	return
	// }
	// res.Success(http.StatusOK, "the user successfully deleted", "", 1)
	c.Status(http.StatusOK)
}
