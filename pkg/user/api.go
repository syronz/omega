package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omega/engine"
	"omega/internal/response"
	"strconv"
)

type API struct {
	Service Service
	Engine  engine.Engine
}

func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

func (p *API) FindAll(c *gin.Context) {
	res := response.Response{Context: c}
	users, err := p.Service.FindAll()
	if err != nil {
		res.Failed(http.StatusExpectationFailed, 1424, "Failed to fetch users data", err)
		return
	}
	res.Success(http.StatusOK, "", users, 200)
	return
}

func (p *API) FindByID(c *gin.Context) {
	res := response.Response{Context: c}
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := p.Service.FindByID(uint(id))

	if err != nil {
		res.Failed(http.StatusExpectationFailed, 1424, "Failed to fetch user data", err)
		return
	}
	res.Success(http.StatusOK, "", user, 200)
	return
}

func (p *API) Create(c *gin.Context) {
	res := response.Response{Context: c}
	var user User

	err := c.BindJSON(&user)
	if err != nil {
		res.Failed(http.StatusBadRequest, 1400, "missing information", "")
		return
	}

	createdUser, err := p.Service.Save(user)
	res.Success(http.StatusOK, "", createdUser, 200)
}

func (p *API) Update(c *gin.Context) {
	res := response.Response{Context: c}

	var user User
	err := c.BindJSON(&user)

	if err != nil {
		res.Failed(http.StatusBadRequest, 1400, "missing update information", "")
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	currentUser, findErr := p.Service.FindByID(uint(id))
	if currentUser.ID == 0 {
		res.Failed(http.StatusBadRequest, 1400, findErr.Error(), "")
		return
	}

	user.ID = currentUser.ID
	updatedUser, updateErr := p.Service.Save(user)
	if updateErr != nil {
		res.Failed(http.StatusBadRequest, 1400, updateErr.Error(), "")
		return
	}
	res.Success(http.StatusOK, string(updatedUser.ID), updatedUser, 200)
	return
}

func (p *API) Delete(c *gin.Context) {
	res := response.Response{Context: c}
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := p.Service.FindByID(uint(id))
	if err != nil {
		res.Failed(http.StatusBadRequest, 1400, "user does not exist", "")
		return
	}

	err = p.Service.Delete(user)
	if err != nil {
		res.Failed(http.StatusBadRequest, 1400, "something went wrong, cannot delete this user", "")
		return
	}
	res.Success(http.StatusOK, "the user successfully deleted", "", 1)
	c.Status(http.StatusOK)
}
