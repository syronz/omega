package role

import (
	"net/http"
	"omega/engine"
	"omega/internal/param"
	"omega/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API for injecting role service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI for role is used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// FindAll roles
func (p *API) FindAll(c *gin.Context) {
	roles, err := p.Service.FindAll()

	if err != nil {
		response.RecordNotFound(c, err, "roles")
		return
	}

	response.Success(c, roles)
}

// List of roles
func (p *API) List(c *gin.Context) {
	params := param.Get(c)

	p.Engine.Debug(params)
	data, err := p.Service.List(params)
	if err != nil {
		response.RecordNotFound(c, err, "roles")
		return
	}

	p.Engine.Record(c, "role-list", params)
	response.Success(c, data)
}

// FindByID is used for fetch a role by his id
func (p *API) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	role, err := p.Service.FindByID(id)

	if err != nil {
		response.RecordNotFound(c, err, "role")
		return
	}

	response.Success(c, role)
}

// Create role
func (p *API) Create(c *gin.Context) {
	var role Role

	err := c.BindJSON(&role)
	if err != nil {
		response.ErrorInBinding(c, err, "create role")
		return
	}

	createdRole, err := p.Service.Save(role)
	if err != nil {
		response.ErrorOnSave(c, err, "role")
		return
	}

	p.Engine.Record(c, "role-create", nil, role)
	response.SuccessSave(c, createdRole, "role/create")
}

// Update role
func (p *API) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var role Role

	if err = c.BindJSON(&role); err != nil {
		response.ErrorInBinding(c, err, "update role")
		return
	}
	role.ID = id

	roleBefore, err := p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "update role")
		return
	}

	updatedRole, err := p.Service.Save(role)
	if err != nil {
		response.ErrorOnSave(c, err, "update role")
		return
	}

	p.Engine.Record(c, "role-update", roleBefore, updatedRole)
	response.SuccessSave(c, updatedRole, "role updated")
}

// Delete role
func (p *API) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var role Role

	role, err = p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "delete role")
		return
	}

	err = p.Service.Delete(role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Something went wrong, cannot delete this role",
			Code:    1500,
		})
		return
	}

	p.Engine.Record(c, "role-delete", role)
	response.SuccessSave(c, role, "role successfully deleted")
}
