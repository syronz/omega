package eacapi

import (
	"net/http"
	"omega/domain/eaccounting"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacterm"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// CurrencyAPI for injecting currency service
type CurrencyAPI struct {
	Service service.EacCurrencyServ
	Engine  *core.Engine
}

// ProvideCurrencyAPI for currency is used in wire
func ProvideCurrencyAPI(c service.EacCurrencyServ) CurrencyAPI {
	return CurrencyAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a currency by it's id
func (p *CurrencyAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var err error
	var currency eacmodel.Currency
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("currencyID"), "E1484534", eacterm.Currency); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if currency, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.ViewCurrency)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, eacterm.Currency).
		JSON(currency)
}

// List of currencies
func (p *CurrencyAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, eacmodel.CurrencyTable, eaccounting.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1464860"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.ListCurrency)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, eacterm.Currencies).
		JSON(data)
}

// Create currency
func (p *CurrencyAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var currency, createdCurrency eacmodel.Currency
	var err error

	if currency.CompanyID, currency.NodeID, err = resp.GetCompanyNode("E1490109", eaccounting.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if currency.CompanyID, err = resp.GetCompanyID("E1490109"); err != nil {
		return
	}

	if !resp.CheckRange(currency.CompanyID) {
		return
	}

	if err = resp.Bind(&currency, "E1412634", eaccounting.Domain, eacterm.Currency); err != nil {
		return
	}

	if createdCurrency, err = p.Service.Create(currency); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(eaccounting.CreateCurrency, currency)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, eacterm.Currency).
		JSON(createdCurrency)
}

// Update currency
func (p *CurrencyAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var err error

	var currency, currencyBefore, currencyUpdated eacmodel.Currency
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("currencyID"), "E1487831", eacterm.Currency); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if err = resp.Bind(&currency, "E1487442", eaccounting.Domain, eacterm.Currency); err != nil {
		return
	}

	if currencyBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	currency.ID = fix.ID
	currency.CompanyID = fix.CompanyID
	currency.NodeID = fix.NodeID
	if currencyUpdated, err = p.Service.Save(currency); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.UpdateCurrency, currencyBefore, currency)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, eacterm.Currency).
		JSON(currencyUpdated)
}

// Delete currency
func (p *CurrencyAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var err error
	var currency eacmodel.Currency
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("currencyID"), "E1477642", eacterm.Currency); err != nil {
		return
	}

	if currency, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.DeleteCurrency, currency)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, eacterm.Currency).
		JSON()
}

// Excel generate excel files eaced on search
func (p *CurrencyAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, eacterm.Currencies, eaccounting.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1435727"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	currencies, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("currency")
	ex.AddSheet("Currencies").
		AddSheet("Summary").
		Active("Currencies").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "F", 15.3).
		SetColWidth("G", "G", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Currencies").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Symbol", "Code", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Symbol", "Code", "UpdatedAt").
		WriteData(currencies).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.ExcelCurrency)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
