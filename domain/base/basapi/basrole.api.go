package basapi

import (
	"net/http"
	"omega/domain/base/basevent"
	"omega/domain/base/basmodel"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/response"
	"omega/internal/term"
	"omega/internal/types"
	"omega/utils/excel"

	"github.com/gin-gonic/gin"
)

const thisBasRole = "role"
const thisBasRoles = "bas_roles"

// BasRoleAPI for injecting role service
type BasRoleAPI struct {
	Service service.BasRoleServ
	Engine  *core.Engine
}

// ProvideBasRoleAPI for role is used in wire
func ProvideBasRoleAPI(c service.BasRoleServ) BasRoleAPI {
	return BasRoleAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a role by it's id
func (p *BasRoleAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var role basmodel.BasRole

	if role.ID, err = types.StrToRowID(c.Param("roleID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if role, err = p.Service.FindByID(role.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(err).MessageT(term.Record_Not_Found).JSON()
		return
	}

	resp.Record(basevent.BasRoleView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisBasRole).
		JSON(role)
}

// List of roles
func (p *BasRoleAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisBasRoles)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasRoleList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisBasRoles).
		JSON(data)
}

// Create role
func (p *BasRoleAPI) Create(c *gin.Context) {
	var role basmodel.BasRole
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisBasRoles)

	createdBasRole, err := p.Service.Create(role, params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasRoleCreate, nil, role)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisBasRole).
		JSON(createdBasRole)
}

// Update role
func (p *BasRoleAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	var role, roleBefore, roleUpdated basmodel.BasRole

	role.ID, err = types.StrToRowID(c.Param("roleID"))
	if err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if roleBefore, err = p.Service.FindByID(role.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if roleUpdated, err = p.Service.Save(role); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasRoleUpdate, roleBefore, role)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisBasRole).
		JSON(roleUpdated)
}

// Delete role
func (p *BasRoleAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var role basmodel.BasRole

	if role.ID, err = types.StrToRowID(c.Param("roleID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisBasRoles)

	if role, err = p.Service.Delete(role.ID, params); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	resp.Record(basevent.BasRoleDelete, role)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisBasRole).
		JSON()
}

// Excel generate excel files based on search
func (p *BasRoleAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisBasRoles)
	roles, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("role")
	ex.AddSheet("BasRoles").
		AddSheet("Summary").
		Active("BasRoles").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "B", 15.3).
		SetColWidth("C", "C", 80).
		SetColWidth("D", "E", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("BasRoles").
		WriteHeader("ID", "Name", "Resources", "Description", "Updated At").
		SetSheetFields("ID", "Name", "Resources", "Description", "UpdatedAt").
		WriteData(roles).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	resp.Record(basevent.BasRoleExcel)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
