package company

import (
	"net/http"
	"omega/engine"
	"omega/internal/param"
	"omega/internal/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API for injecting company service
type API struct {
	Service Service
	Engine  engine.Engine
}

// ProvideAPI for company is used in wire
func ProvideAPI(p Service) API {
	return API{Service: p, Engine: p.Engine}
}

// FindAll companies
func (p *API) FindAll(c *gin.Context) {
	if p.Engine.CheckAccess(c, "companies:names") {
		response.NoPermission(c)
		return
	}
	companies, err := p.Service.FindAll()

	if err != nil {
		response.RecordNotFound(c, err, "companies")
		return
	}

	response.Success(c, companies)
}

// List of companies
func (p *API) List(c *gin.Context) {
	if p.Engine.CheckAccess(c, "companies:read") {
		response.NoPermissionRecord(c, p.Engine, "company-list-forbidden")
		return
	}

	params := param.Get(c)

	p.Engine.Debug(params)
	data, err := p.Service.List(params)
	if err != nil {
		response.RecordNotFound(c, err, "companies")
		return
	}

	p.Engine.Record(c, "company-list")
	response.Success(c, data)
}

// FindByID is used for fetch a company by his id
func (p *API) FindByID(c *gin.Context) {
	if p.Engine.CheckAccess(c, "companies:read") {
		response.NoPermissionRecord(c, p.Engine, "company-view-forbidden")
		return
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	company, err := p.Service.FindByID(id)

	if err != nil {
		response.RecordNotFound(c, err, "company")
		return
	}

	p.Engine.Record(c, "company-view")
	response.Success(c, company)
}

// Create company
func (p *API) Create(c *gin.Context) {
	var company Company

	err := c.BindJSON(&company)
	if err != nil {
		response.ErrorInBinding(c, err, "create company")
		return
	}

	if p.Engine.CheckAccess(c, "companies:write") {
		response.NoPermissionRecord(c, p.Engine, "company-create-forbidden", nil, company)
		return
	}

	createdCompany, err := p.Service.Save(company)
	if err != nil {
		response.ErrorOnSave(c, err, "company")
		return
	}

	p.Engine.Record(c, "company-create", nil, company)
	response.SuccessSave(c, createdCompany, "company/create")
}

// Update company
func (p *API) Update(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}

	var company Company

	if err = c.BindJSON(&company); err != nil {
		response.ErrorInBinding(c, err, "update company")
		return
	}
	company.ID = id
	if p.Engine.CheckAccess(c, "companies:write") {
		response.NoPermissionRecord(c, p.Engine, "company-update-forbidden", nil, company)
		return
	}

	companyBefore, err := p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "update company")
		return
	}

	updatedCompany, err := p.Service.Save(company)
	if err != nil {
		response.ErrorOnSave(c, err, "update company")
		return
	}

	p.Engine.Record(c, "company-update", companyBefore, updatedCompany)
	response.SuccessSave(c, updatedCompany, "company updated")
}

// Delete company
func (p *API) Delete(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 16)
	if err != nil {
		response.InvalidID(c, err)
		return
	}
	if p.Engine.CheckAccess(c, "companies:write") {
		response.NoPermissionRecord(c, p.Engine, "company-delete-forbidden", nil, id)
		return
	}

	var company Company

	company, err = p.Service.FindByID(id)
	if err != nil {
		response.RecordNotFound(c, err, "delete company")
		return
	}

	err = p.Service.Delete(company)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Something went wrong, cannot delete this company",
			Code:    1500,
		})
		return
	}

	p.Engine.Record(c, "company-delete", company)
	response.SuccessSave(c, company, "company successfully deleted")
}
