package eacapi

import (
	"net/http"
	"omega/domain/eaccounting"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacterm"
	"omega/domain/eaccounting/enum/transactiontype"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// TransactionAPI for injecting transaction service
type TransactionAPI struct {
	Service service.EacTransactionServ
	Engine  *core.Engine
}

// ProvideTransactionAPI for transaction is used in wire
func ProvideTransactionAPI(c service.EacTransactionServ) TransactionAPI {
	return TransactionAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a transaction by it's id
func (p *TransactionAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var err error
	var transaction eacmodel.Transaction
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("transactionID"), "E1423147", eacterm.Transaction); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if transaction, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.ViewTransaction)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, eacterm.Transaction).
		JSON(transaction)
}

// List of transactions
func (p *TransactionAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, eacmodel.TransactionTable, eaccounting.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1446041"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.ListTransaction)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, eacterm.Transactions).
		JSON(data)
}

// ManualTransfer transaction
func (p *TransactionAPI) ManualTransfer(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, eacterm.Transactions, eaccounting.Domain)
	var transaction, createdTransaction eacmodel.Transaction
	var err error

	if transaction.CompanyID, transaction.NodeID, err = resp.GetCompanyNode("E1495400", eaccounting.Domain); err != nil {
		resp.Error(err).JSON()
		return
	}

	if transaction.CompanyID, err = resp.GetCompanyID("E1483718"); err != nil {
		return
	}

	if !resp.CheckRange(transaction.CompanyID, transaction.NodeID) {
		return
	}

	if err = resp.Bind(&transaction, "E1444992", eaccounting.Domain, eacterm.Transaction); err != nil {
		return
	}

	transaction.Type = transactiontype.Manual
	transaction.CreatedBy = params.UserID

	if createdTransaction, err = p.Service.Transfer(transaction); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(eaccounting.ManualTransfer, transaction)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, eacterm.Transaction).
		JSON(createdTransaction)
}

// Update transaction
func (p *TransactionAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var err error

	var transaction, transactionBefore, transactionUpdated eacmodel.Transaction
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("transactionID"), "E1452724", eacterm.Transaction); err != nil {
		return
	}

	if !resp.CheckRange(fix.CompanyID, fix.NodeID) {
		return
	}

	if err = resp.Bind(&transaction, "E1451486", eaccounting.Domain, eacterm.Transaction); err != nil {
		return
	}

	if transactionBefore, err = p.Service.FindByID(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	transaction.CreatedBy = transactionBefore.CreatedBy
	transaction.Hash = transactionBefore.Hash
	transaction.Type = transactionBefore.Type
	transaction.ID = fix.ID
	transaction.CompanyID = fix.CompanyID
	transaction.NodeID = fix.NodeID
	if transactionUpdated, err = p.Service.Save(transaction); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.UpdateTransaction, transactionBefore, transaction)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, eacterm.Transaction).
		JSON(transactionUpdated)
}

// Delete transaction
func (p *TransactionAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, eaccounting.Domain)
	var err error
	var transaction eacmodel.Transaction
	var fix types.FixedCol

	if fix, err = resp.GetFixedCol(c.Param("transactionID"), "E1413663", eacterm.Transaction); err != nil {
		return
	}

	if transaction, err = p.Service.Delete(fix); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.DeleteTransaction, transaction)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, eacterm.Transaction).
		JSON()
}

// Excel generate excel files eaced on search
func (p *TransactionAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, eacterm.Transactions, eaccounting.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E1469354"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	transactions, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("transaction")
	ex.AddSheet("Transactions").
		AddSheet("Summary").
		Active("Transactions").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "F", 15.3).
		SetColWidth("G", "G", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Transactions").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Symbol", "Code", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Symbol", "Code", "UpdatedAt").
		WriteData(transactions).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(eaccounting.ExcelTransaction)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
