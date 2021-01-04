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

// AccountAPI for injecting account service
type AccountAPI struct {
	Service service.BasAccountServ
	Engine  *core.Engine
}

// ProvideAccountAPI for account is used in wire
func ProvideAccountAPI(c service.BasAccountServ) AccountAPI {
	return AccountAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a account by it's id
func (p *AccountAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var account basmodel.Account
	var fix types.FixedNode

	if fix, err = resp.GetFixedNode(c.Param("accountID"), "E1070061", basterm.Account); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if account, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ViewAccount)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.Account).
		JSON(account)
}

// List of accounts
func (p *AccountAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basmodel.AccountTable, base.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1056290"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ListAccount)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Accounts).
		JSON(data)
}

// Create account
func (p *AccountAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var account, createdAccount basmodel.Account
	var err error

	if account.CompanyID, account.NodeID, err = resp.GetCompanyNode("E1057239", base.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if account.CompanyID, err = resp.GetCompanyID("E1085677"); err != nil {
		return
	}

	if !resp.CheckRange(account.CompanyID) {
		return
	}

	if err = resp.Bind(&account, "E1057541", base.Domain, basterm.Account); err != nil {
		return
	}

	if createdAccount, err = p.Service.Create(account); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(base.CreateAccount, account)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, basterm.Account).
		JSON(createdAccount)
}

// Update account
func (p *AccountAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error

	var account, accountBefore, accountUpdated basmodel.Account
	var fix types.FixedNode

	if fix, err = resp.GetFixedNode(c.Param("accountID"), "E1076703", basterm.Account); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if err = resp.Bind(&account, "E1086162", base.Domain, basterm.Account); err != nil {
		return
	}

	if accountBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	account.ID = fix.ID
	account.CompanyID = fix.CompanyID
	account.NodeID = fix.NodeID
	if accountUpdated, err = p.Service.Save(account); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.UpdateAccount, accountBefore, account)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Account).
		JSON(accountUpdated)
}

// Delete account
func (p *AccountAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, base.Domain)
	var err error
	var account basmodel.Account
	var fix types.FixedNode

	if fix, err = resp.GetFixedNode(c.Param("accountID"), "E1092196", basterm.Account); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID) {
		return
	}

	if account, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.DeleteAccount, account)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, basterm.Account).
		JSON()
}

// Excel generate excel files based on search
func (p *AccountAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Accounts, base.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1066535"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID) {
		return
	}

	accounts, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("account")
	ex.AddSheet("Accounts").
		AddSheet("Summary").
		Active("Accounts").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "G", 15.3).
		SetColWidth("H", "H", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Accounts").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Code", "Type", "Status", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Code", "Type", "Status", "UpdatedAt").
		WriteData(accounts).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(base.ExcelAccount)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
