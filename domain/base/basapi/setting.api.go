package basapi

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corerr"
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

	if setting.ID, err = types.StrToRowID(c.Param("settingID")); err != nil {
		resp.Status(http.StatusNotAcceptable).Error(err).MessageT(corerr.InvalidID).JSON()
		return
	}

	if setting, err = p.Service.FindByID(setting.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ViewSetting)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.Setting).
		JSON(setting)
}

// FindByProperty is used when we try to find a setting with property
func (p *SettingAPI) FindByProperty(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	property := c.Param("property")

	setting, err := p.Service.FindByProperty(property)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(err).JSON()
		return
	}

	resp.Status(http.StatusOK).JSON(setting)
}

// List of settings
func (p *SettingAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.RoleTable, base.Domain)

	data := make(map[string]interface{})
	var err error

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

	if setting.ID, err = resp.GetRowID(c.Param("settingID"), "E1074247",
		basterm.Setting); err != nil {
		return
	}

	if err = resp.Bind(&setting, "E1049049", base.Domain, basterm.Setting); err != nil {
		return
	}

	if settingBefore, err = p.Service.FindByID(setting.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	if settingUpdated, err = p.Service.Update(setting); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.UpdateSetting, settingBefore, settingUpdated)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Setting).
		JSON(settingUpdated)
}

// Delete setting
func (p *SettingAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var setting basmodel.Setting

	if setting.ID, err = resp.GetRowID(c.Param("settingID"), "E1076780",
		basterm.Setting); err != nil {
		return
	}

	if setting, err = p.Service.Delete(setting.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.DeleteSetting, setting)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, basterm.Setting).
		JSON()
}

// Excel generate excel files based on search
func (p *SettingAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Roles, base.Domain)

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
		SetColWidth("A", "A", 20).
		SetColWidth("B", "C", 15.3).
		SetColWidth("F", "F", 20).
		SetColWidth("L", "M", 20).
		Active("Summary").
		Active("Settings").
		WriteHeader("ID", "Property", "Value", "Type", "Description", "Created At", "Updated At").
		SetSheetFields("ID", "Property", "Value", "Type", "Description", "CreatedAt", "UpdatedAt").
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
