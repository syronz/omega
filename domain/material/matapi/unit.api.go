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

// UnitAPI for injecting unit service
type UnitAPI struct {
	Service service.MatUnitServ
	Engine  *core.Engine
}

// ProvideUnitAPI for unit is used in wire
func ProvideUnitAPI(c service.MatUnitServ) UnitAPI {
	return UnitAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a unit by it's id
func (p *UnitAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error
	var unit matmodel.Unit
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("unitID"), "E7121585", matterm.Unit); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if unit, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ViewUnit)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, matterm.Unit).
		JSON(unit)
}

// List of units
func (p *UnitAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, matmodel.UnitTable, material.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E7150309"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ListUnit)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, matterm.Units).
		JSON(data)
}

// Create unit
func (p *UnitAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var unit, createdUnit matmodel.Unit
	var err error

	if unit.CompanyID, unit.NodeID, err = resp.GetCompanyNode("E7131572", material.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if unit.CompanyID, err = resp.GetCompanyID("E7156949"); err != nil {
		return
	}

	if !resp.CheckRange(unit.CompanyID, unit.NodeID) {
		return
	}

	if err = resp.Bind(&unit, "E7134015", material.Domain, matterm.Unit); err != nil {
		return
	}

	if createdUnit, err = p.Service.Create(unit); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(material.CreateUnit, unit)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, matterm.Unit).
		JSON(createdUnit)
}

// Update unit
func (p *UnitAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error

	var unit, unitBefore, unitUpdated matmodel.Unit
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("unitID"), "E7155107", matterm.Unit); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if err = resp.Bind(&unit, "E7191090", material.Domain, matterm.Unit); err != nil {
		return
	}

	if unitBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	unit.ID = fix.ID
	unit.CompanyID = fix.CompanyID
	unit.NodeID = fix.NodeID
	if unitUpdated, err = p.Service.Save(unit); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.UpdateUnit, unitBefore, unit)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, matterm.Unit).
		JSON(unitUpdated)
}

// Delete unit
func (p *UnitAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, material.Domain)
	var err error
	var unit matmodel.Unit
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("unitID"), "E7146958", matterm.Unit); err != nil {
		return
	}

	if unit, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.DeleteUnit, unit)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, matterm.Unit).
		JSON()
}

// Excel generate excel files eaced on search
func (p *UnitAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, matterm.Units, material.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E7159543"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	units, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	type ExUnit struct {
		matmodel.Unit
		ExDescription string `json:"ex_description"`
	}

	exUnits := make([]ExUnit, len(units))

	for i, v := range units {
		exUnits[i].ID = v.ID
		exUnits[i].CompanyID = v.CompanyID
		exUnits[i].NodeID = v.NodeID
		exUnits[i].Name = v.Name
		if v.Description != nil {
			exUnits[i].ExDescription = *v.Description
		}
		exUnits[i].UpdatedAt = v.UpdatedAt
	}

	ex := excel.New("unit")
	ex.AddSheet("Units").
		AddSheet("Summary").
		Active("Units").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "E", 15.3).
		SetColWidth("F", "F", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Units").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Description", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "ExDescription", "UpdatedAt").
		WriteData(exUnits).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(material.ExcelUnit)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
