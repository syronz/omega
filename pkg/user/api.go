package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omega/engine"
	"omega/internal/response"
	"strconv"
)

// API for injecting service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// FindAll users
func (p *API) FindAll(c *gin.Context) {
	users, err := p.Service.FindAll()

	if err != nil {
		return
	}

	response.SuccessAll(c, users)
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
		response.RecordNotFound(c, err, "User")
		return
	}

	response.SuccessOne(c, user)
}

// Create user
func (p *API) Create(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.ErrorInBinding(c, err, "Create User")
		return
	}

	createdUser, err := p.Service.Save(user)
	if err != nil {
		response.ErrorInCreating(c, err, "User")
		return
	}

	response.SuccessOne(c, createdUser)
}

// Update user
func (p *API) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var user User

	// currentUser, findErr := p.Service.FindByID(uint(id))
	// if currentUser.ID == 0 {
	// 	res.Failed(http.StatusBadRequest, 1400, findErr.Error(), "")
	// 	return
	// }

	err = c.BindJSON(&user)
	if err != nil {
		response.ErrorInBinding(c, err, " Update User")
		return
	}
	user.ID = id

	createdUser, err := p.Service.Save(user)
	if err != nil {
		response.ErrorInCreating(c, err, " Update User")
		return
	}

	_, updateErr := p.Service.Save(user)
	if updateErr != nil {
		// res.Failed(http.StatusBadRequest, 1400, updateErr.Error(), "")
		return
	}

	response.SuccessOne(c, createdUser)

	// if err != nil {
	// 	res.Failed(http.StatusBadRequest, 1400, "missing update information", "")
	// 	return
	// }

	// res.Success(http.StatusOK, string(updatedUser.ID), updatedUser, 200)
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
