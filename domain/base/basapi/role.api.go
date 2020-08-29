package basapi

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/response"
	"omega/internal/term"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

const thisRole = "role"
const thisRoles = "bas_roles"

// RoleAPI for injecting role service
type RoleAPI struct {
	Service service.BasRoleServ
	Engine  *core.Engine
}

// ProvideRoleAPI for role is used in wire
func ProvideRoleAPI(c service.BasRoleServ) RoleAPI {
	return RoleAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a role by it's id
func (p *RoleAPI) FindByID(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, thisRoles)
	var err error
	var role basmodel.Role

	if role.ID, err = types.StrToRowID(c.Param("roleID")); err != nil {
		resp.NotBind("E1015984", base.Domain, err, "ID", "/roles/:roleID")
		return
	}

	if role, err = p.Service.FindByID(params, role.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ViewRole)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisRole).
		JSON(role)
}

// List of roles
func (p *RoleAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, thisRoles)

	data := make(map[string]interface{})
	var err error

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ListRole)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisRoles).
		JSON(data)
}

// Create role
func (p *RoleAPI) Create(c *gin.Context) {
	var role basmodel.Role
	resp := response.New(p.Engine, c)

	if err := c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisRoles)

	createdRole, err := p.Service.Create(role, params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.CreateRole, nil, role)

	resp.Status(http.StatusOK).
		MessageT(term.V_created_successfully, thisRole).
		JSON(createdRole)
}

// Update role
func (p *RoleAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	var role, roleBefore, roleUpdated basmodel.Role

	role.ID, err = types.StrToRowID(c.Param("roleID"))
	if err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&role); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	params := param.Get(c, p.Engine, thisRoles)

	if roleBefore, err = p.Service.FindByID(params, role.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if roleUpdated, err = p.Service.Save(role); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.UpdateRole, roleBefore, role)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisRole).
		JSON(roleUpdated)
}

// Delete role
func (p *RoleAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var role basmodel.Role

	if role.ID, err = types.StrToRowID(c.Param("roleID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	params := param.Get(c, p.Engine, thisRoles)

	if role, err = p.Service.Delete(role.ID, params); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	resp.Record(base.DeleteRole, role)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisRole).
		JSON()
}

// Excel generate excel files based on search
func (p *RoleAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisRoles)
	roles, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("role")
	ex.AddSheet("Roles").
		AddSheet("Summary").
		Active("Roles").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "B", 15.3).
		SetColWidth("C", "C", 80).
		SetColWidth("D", "E", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Roles").
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

	resp.Record(base.ExcelRole)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
