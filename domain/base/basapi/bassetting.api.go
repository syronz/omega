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

const thisBasSetting = "setting"
const thisBasSettings = "settings"

// BasSettingAPI for injecting setting service
type BasSettingAPI struct {
	Service service.BasSettingServ
	Engine  *core.Engine
}

// ProvideBasSettingAPI for setting is used in wire
func ProvideBasSettingAPI(c service.BasSettingServ) BasSettingAPI {
	return BasSettingAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a setting by it's id
func (p *BasSettingAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var setting basmodel.BasSetting

	if setting.ID, err = types.StrToRowID(c.Param("settingID")); err != nil {
		resp.Status(http.StatusNotAcceptable).Error(err).MessageT(term.Invalid_ID).JSON()
		return
	}

	if setting, err = p.Service.FindByID(setting.ID); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasSettingView)
	resp.Status(http.StatusOK).
		MessageT(term.V_info, thisBasSetting).
		JSON(setting)
}

// FindByProperty is used when we try to find a setting with property
func (p *BasSettingAPI) FindByProperty(c *gin.Context) {
	resp := response.New(p.Engine, c)
	property := c.Param("property")

	setting, err := p.Service.FindByProperty(property)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(err).JSON()
		return
	}

	resp.Status(http.StatusOK).JSON(setting)
}

// List of settings
func (p *BasSettingAPI) List(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisBasSettings)

	data, err := p.Service.List(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasSettingList)
	resp.Status(http.StatusOK).
		MessageT(term.List_of_V, thisBasSettings).
		JSON(data)
}

// Update setting
func (p *BasSettingAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error

	var setting, settingBefore, settingUpdated basmodel.BasSetting

	if setting.ID, err = types.StrToRowID(c.Param("settingID")); err != nil {
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if err = c.ShouldBindJSON(&setting); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, err)
		return
	}

	if settingBefore, err = p.Service.FindByID(setting.ID); err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	if settingUpdated, err = p.Service.Update(setting); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(basevent.BasSettingUpdate, settingBefore, settingUpdated)

	resp.Status(http.StatusOK).
		MessageT(term.V_updated_successfully, thisBasSetting).
		JSON(settingUpdated)

}

// Delete setting
func (p *BasSettingAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c)
	var err error
	var setting basmodel.BasSetting

	if setting.ID, err = types.StrToRowID(c.Param("settingID")); err != nil {
		p.Engine.CheckError(err, err.Error())
		resp.Error(term.Invalid_ID).JSON()
		return
	}

	if setting, err = p.Service.Delete(setting.ID); err != nil {
		resp.Status(http.StatusInternalServerError).Error(err).JSON()
		return
	}

	resp.Record(basevent.BasSettingDelete, setting)
	resp.Status(http.StatusOK).
		MessageT(term.V_deleted_successfully, thisBasSetting).
		JSON()
}

// Excel generate excel files based on search
func (p *BasSettingAPI) Excel(c *gin.Context) {
	resp := response.New(p.Engine, c)

	params := param.Get(c, p.Engine, thisBasSettings)
	settings, err := p.Service.Excel(params)
	if err != nil {
		resp.Status(http.StatusNotFound).Error(term.Record_Not_Found).JSON()
		return
	}

	ex := excel.New("node").
		AddSheet("Nodes").
		AddSheet("Summary").
		Active("Nodes").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("A", "A", 20).
		SetColWidth("B", "C", 15.3).
		SetColWidth("F", "F", 20).
		SetColWidth("L", "M", 20).
		Active("Summary").
		Active("Nodes").
		WriteHeader("ID", "Name", "BasSettingname", "Code", "Status", "Role",
			"Language", "Type", "Email", "Readonly", "Direction", "Created At",
			"Updated At").
		SetSheetFields("ID", "Name", "LegalName", "ServerAddress", "Expiration", "Plan",
			"Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
		WriteData(settings).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &response.Result{
			Message: "Error in generating Excel file",
		})
		return
	}

	resp.Record(basevent.BasSettingExcel)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
