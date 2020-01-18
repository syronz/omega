package user

import (
	"net/http"
	"omega/engine"
	"omega/internal/param"
	"omega/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
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

// List of ousers
func (p *API) List(c *gin.Context) {
	params := param.Get(c)

	p.Engine.Debug(params)

	response.SuccessMessage(c, "it works fine")

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

	response.SuccessSave(c, createdUser, "user created")
}

// Update user
func (p *API) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var user User

	if err = c.BindJSON(&user); err != nil {
		response.ErrorInBinding(c, err, "update user")
		return
	}
	user.ID = id

	updatedUser, err := p.Service.Save(user)
	if err != nil {
		response.ErrorOnSave(c, err, "update user")
		return
	}

	response.SuccessSave(c, updatedUser, "user updated")
}

// Delete user
func (p *API) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var user User

	user, err = p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "delete user")
		return
	}

	err = p.Service.Delete(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Something went wrong, cannot delete this user",
			Code:    1500,
		})
		return
	}

	c.Status(http.StatusOK)
	response.SuccessSave(c, user, "user successfully deleted")
}
