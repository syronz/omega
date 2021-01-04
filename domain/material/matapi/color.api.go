package matapi

import (
	"net/http"
	"omega/domain/base/message/basterm"
	"omega/domain/material"
	"omega/domain/material/matmodel"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// ColorAPI for injecting color service
type ColorAPI struct {
	Service service.MatColorServ
	Engine  *core.Engine
}

// ProvideColorAPI for color is used in wire
func ProvideColorAPI(c service.MatColorServ) ColorAPI {
	return ColorAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a color by it's id
func (p *ColorAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error
	var color matmodel.Color
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("colorID"), "E7177921", basterm.Color); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if color, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ViewColor)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.Color).
		JSON(color)
}

// List of colors
func (p *ColorAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, matmodel.ColorTable, material.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E7137844"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ListColor)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Colors).
		JSON(data)
}

// Create color
func (p *ColorAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var color, createdColor matmodel.Color
	var err error

	if color.CompanyID, color.NodeID, err = resp.GetCompanyNode("E7163888", material.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if color.CompanyID, err = resp.GetCompanyID("E7193880"); err != nil {
		return
	}

	if !resp.CheckRange(color.CompanyID) {
		return
	}

	if err = resp.Bind(&color, "E7164911", material.Domain, basterm.Color); err != nil {
		return
	}

	if createdColor, err = p.Service.Create(color); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(material.CreateColor, color)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, basterm.Color).
		JSON(createdColor)
}

// Update color
func (p *ColorAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error

	var color, colorBefore, colorUpdated matmodel.Color
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("colorID"), "E7110043", basterm.Color); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if err = resp.Bind(&color, "E7187070", material.Domain, basterm.Color); err != nil {
		return
	}

	if colorBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	color.ID = fix.ID
	color.CompanyID = fix.CompanyID
	color.NodeID = fix.NodeID
	if colorUpdated, err = p.Service.Save(color); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.UpdateColor, colorBefore, color)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Color).
		JSON(colorUpdated)
}

// Delete color
func (p *ColorAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error
	var color matmodel.Color
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("colorID"), "E7111561", basterm.Color); err != nil {
		return
	}

	if color, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.DeleteColor, color)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, basterm.Color).
		JSON()
}

// Excel generate excel files eaced on search
func (p *ColorAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Colors, material.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E7133217"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	colors, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("color")
	ex.AddSheet("Colors").
		AddSheet("Summary").
		Active("Colors").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "F", 15.3).
		SetColWidth("G", "G", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Colors").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Symbol", "Code", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Symbol", "Code", "UpdatedAt").
		WriteData(colors).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ExcelColor)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
