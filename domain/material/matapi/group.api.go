package matapi

import (
	"net/http"
	"omega/domain/material"
	"omega/domain/material/matmodel"
	"omega/domain/material/matterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// GroupAPI for injecting group service
type GroupAPI struct {
	Service service.MatGroupServ
	Engine  *core.Engine
}

// ProvideGroupAPI for group is used in wire
func ProvideGroupAPI(c service.MatGroupServ) GroupAPI {
	return GroupAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a group by it's id
func (p *GroupAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error
	var group matmodel.Group
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("groupID"), "E7190412", matterm.Group); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if group, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ViewGroup)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, matterm.Group).
		JSON(group)
}

// List of groups
func (p *GroupAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, matmodel.GroupTable, material.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E7158770"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ListGroup)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, matterm.Groups).
		JSON(data)
}

// Create group
func (p *GroupAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var group, createdGroup matmodel.Group
	var err error

	if group.CompanyID, group.NodeID, err = resp.GetCompanyNode("E7142750", material.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if group.CompanyID, err = resp.GetCompanyID("E7195736"); err != nil {
		return
	}

	if !resp.CheckRange(group.CompanyID, group.NodeID) {
		return
	}

	if err = resp.Bind(&group, "E7149387", material.Domain, matterm.Group); err != nil {
		return
	}

	if createdGroup, err = p.Service.Create(group); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(material.CreateGroup, group)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, matterm.Group).
		JSON(createdGroup)
}

// Update group
func (p *GroupAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error

	var group, groupBefore, groupUpdated matmodel.Group
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("groupID"), "E7127629", matterm.Group); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if err = resp.Bind(&group, "E7145224", material.Domain, matterm.Group); err != nil {
		return
	}

	if groupBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	group.ID = fix.ID
	group.CompanyID = fix.CompanyID
	group.NodeID = fix.NodeID
	if groupUpdated, err = p.Service.Save(group); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.UpdateGroup, groupBefore, group)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, matterm.Group).
		JSON(groupUpdated)
}

// Delete group
func (p *GroupAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error
	var group matmodel.Group
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("groupID"), "E7181189", matterm.Group); err != nil {
		return
	}

	if group, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.DeleteGroup, group)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, matterm.Group).
		JSON()
}

// Excel generate excel files eaced on search
func (p *GroupAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, matterm.Groups, material.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E7177413"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	groups, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("group")
	ex.AddSheet("Groups").
		AddSheet("Summary").
		Active("Groups").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "F", 15.3).
		SetColWidth("G", "G", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Groups").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Symbol", "Code", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Symbol", "Code", "UpdatedAt").
		WriteData(groups).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ExcelGroup)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
