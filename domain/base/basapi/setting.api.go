package basapi

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// SettingAPI for injecting setting service
type SettingAPI struct {
	Service service.BasSettingServ
	Engine  *core.Engine
}

// ProvideSettingAPI for setting is used in wire
func ProvideSettingAPI(c service.BasSettingServ) SettingAPI {
	return SettingAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a setting by it's id
func (p *SettingAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var setting basmodel.Setting
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("settingID"), "E1013513", basterm.Setting); err != nil {
		return
	}

	accessService := service.ProvideBasAccessService(basrepo.ProvideAccessRepo(p.Engine))
	accessResult := accessService.CheckAccess(c, base.SuperAccess)
	if accessResult == true {
		if !resp.CheckRange(fix.CompanyID) {
			return
		}
	}

	if setting, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ViewSetting)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.Setting).
		JSON(setting)
}

// List of settings
func (p *SettingAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.SettingTable, base.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1078032"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ListSetting)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Settings).
		JSON(data)
}

// Update setting
func (p *SettingAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error

	var setting, settingBefore, settingUpdated basmodel.Setting
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("settingID"), "E1074247", basterm.Setting); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if err = resp.Bind(&setting, "E1049049", base.Domain, basterm.Setting); err != nil {
		return
	}

	if settingBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	setting.ID = fix.ID
	setting.CompanyID = fix.CompanyID
	setting.NodeID = fix.NodeID
	if settingUpdated, err = p.Service.Update(setting); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.UpdateSetting, settingBefore, settingUpdated)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Setting).
		JSON(settingUpdated)
}

// Excel generate excel files based on search
func (p *SettingAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Setting, base.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1072378"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	settings, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("setting").
		AddSheet("Settings").
		AddSheet("Summary").
		Active("Settings").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("A", "C", 12).
		SetColWidth("D", "E", 20).
		SetColWidth("G", "G", 75).
		SetColWidth("H", "I", 30).
		SetColWidth("N", "O", 20).
		Active("Summary").
		Active("Settings").
		WriteHeader("ID", "Company ID", "Node ID", "Property", "Value", "Type", "Description", "Created At", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Property", "Value", "Type", "Description", "CreatedAt", "UpdatedAt").
		WriteData(settings).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ExcelSetting)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
