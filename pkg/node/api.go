package node

import (
	"net/http"
	"omega/engine"
	"omega/internal/param"
	"omega/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API for injecting node service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI for node is used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// FindAll nodes
func (p *API) FindAll(c *gin.Context) {
	if p.Engine.CheckAccess(c, "nodes:names") {
		response.NoPermission(c)
		return
	}
	nodes, err := p.Service.FindAll()

	if err != nil {
		response.RecordNotFound(c, err, "nodes")
		return
	}

	response.Success(c, nodes)
}

// List of nodes
func (p *API) List(c *gin.Context) {
	if p.Engine.CheckAccess(c, "nodes:read") {
		response.NoPermissionRecord(c, p.Engine, "node-list-forbidden")
		return
	}

	params := param.Get(c)

	p.Engine.Debug(params)
	data, err := p.Service.List(params)
	if err != nil {
		response.RecordNotFound(c, err, "nodes")
		return
	}

	p.Engine.Record(c, "node-list")
	response.Success(c, data)
}

// FindByID is used for fetch a node by his id
func (p *API) FindByID(c *gin.Context) {
	if p.Engine.CheckAccess(c, "nodes:read") {
		response.NoPermissionRecord(c, p.Engine, "node-view-forbidden")
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	node, err := p.Service.FindByID(id)

	if err != nil {
		response.RecordNotFound(c, err, "node")
		return
	}

	p.Engine.Record(c, "node-view")
	response.Success(c, node)
}

// Create node
func (p *API) Create(c *gin.Context) {
	var node Node

	err := c.BindJSON(&node)
	if err != nil {
		response.ErrorInBinding(c, err, "create node")
		return
	}

	if p.Engine.CheckAccess(c, "nodes:write") {
		response.NoPermissionRecord(c, p.Engine, "node-create-forbidden", nil, node)
		return
	}

	createdNode, err := p.Service.Save(node)
	if err != nil {
		response.ErrorOnSave(c, err, "node")
		return
	}

	p.Engine.Record(c, "node-create", nil, node)
	response.SuccessSave(c, createdNode, "node/create")
}

// Update node
func (p *API) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var node Node

	if err = c.BindJSON(&node); err != nil {
		response.ErrorInBinding(c, err, "update node")
		return
	}
	node.ID = id
	if p.Engine.CheckAccess(c, "nodes:write") {
		response.NoPermissionRecord(c, p.Engine, "node-update-forbidden", nil, node)
		return
	}

	nodeBefore, err := p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "update node")
		return
	}

	updatedNode, err := p.Service.Save(node)
	if err != nil {
		response.ErrorOnSave(c, err, "update node")
		return
	}

	p.Engine.Record(c, "node-update", nodeBefore, updatedNode)
	response.SuccessSave(c, updatedNode, "node updated")
}

// Delete node
func (p *API) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}
	if p.Engine.CheckAccess(c, "nodes:write") {
		response.NoPermissionRecord(c, p.Engine, "node-delete-forbidden", nil, id)
		return
	}

	var node Node

	node, err = p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "delete node")
		return
	}

	err = p.Service.Delete(node)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Something went wrong, cannot delete this node",
			Code:    1500,
		})
		return
	}

	p.Engine.Record(c, "node-delete", node)
	response.SuccessSave(c, node, "node successfully deleted")
}
