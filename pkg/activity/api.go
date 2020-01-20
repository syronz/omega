package activity

import (
	"omega/engine"
	"omega/internal/param"
	"omega/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API for injecting activity service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI for activity is used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// List of activities
func (p *API) List(c *gin.Context) {
	params := param.Get(c)

	p.Engine.Debug(params)
	data, err := p.Service.List(params)
	if err != nil {
		response.RecordNotFound(c, err, "activities")
		return
	}

	response.Success(c, data)
}

// FindByID is used for fetch a activity by his id
func (p *API) FindByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	activity, err := p.Service.FindByID(id)

	if err != nil {
		response.RecordNotFound(c, err, "activity")
		return
	}

	response.Success(c, activity)
}
