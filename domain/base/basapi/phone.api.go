package basapi

import (
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// PhoneAPI for injecting phone service
type PhoneAPI struct {
	Service service.BasPhoneServ
	Engine  *core.Engine
}

// ProvidePhoneAPI for phone is used in wire
func ProvidePhoneAPI(c service.BasPhoneServ) PhoneAPI {
	return PhoneAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a phone by it's id
func (p *PhoneAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var phone basmodel.Phone
	var fix types.FixedNode

	if fix, err = resp.GetFixedNode(c.Param("phoneID"), "E1030412", basterm.Phone); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if phone, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ViewPhone)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.Phone).
		JSON(phone)
}

// List of phones
func (p *PhoneAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.PhoneTable, base.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1062683"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ListPhone)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Phones).
		JSON(data)
}

// Create phone
func (p *PhoneAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var phone, createdPhone basmodel.Phone
	var err error

	if err = resp.Bind(&phone, "E1053717", base.Domain, basterm.Phone); err != nil {
		return
	}

	if createdPhone, err = p.Service.Create(phone); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(base.CreatePhone, phone)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, basterm.Phone).
		JSON(createdPhone)
}

// Update phone
func (p *PhoneAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error

	var phone, phoneBefore, phoneUpdated basmodel.Phone
	var fix types.FixedNode

	if fix, err = resp.GetFixedNode(c.Param("phoneID"), "E1056260", basterm.Phone); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if err = resp.Bind(&phone, "E1073908", base.Domain, basterm.Phone); err != nil {
		return
	}

	if phoneBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	phone.ID = fix.ID
	if phoneUpdated, err = p.Service.Save(phone); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.UpdatePhone, phoneBefore, phone)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Phone).
		JSON(phoneUpdated)
}

// Delete phone
func (p *PhoneAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var phone basmodel.Phone
	var fix types.FixedNode

	if fix, err = resp.GetFixedNode(c.Param("phoneID"), "E1076860", basterm.Phone); err != nil {
		return
	}

	if phone, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.DeletePhone, phone)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, basterm.Phone).
		JSON()
}

// Excel generate excel files based on search
func (p *PhoneAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Phones, base.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1075215"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	phones, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("phone")
	ex.AddSheet("Phones").
		AddSheet("Summary").
		Active("Phones").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "G", 15.3).
		SetColWidth("H", "H", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Phones").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Code", "Type", "Status", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Code", "Type", "Status", "UpdatedAt").
		WriteData(phones).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ExcelPhone)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
