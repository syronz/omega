package service

import (
	"fmt"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// EacTransactionServ for injecting auth eacrepo
type EacTransactionServ struct {
	Repo   eacrepo.TransactionRepo
	Engine *core.Engine
}

// ProvideEacTransactionService for transaction is used in wire
func ProvideEacTransactionService(p eacrepo.TransactionRepo) EacTransactionServ {
	return EacTransactionServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting transaction by it's id
func (p *EacTransactionServ) FindByID(fix types.FixedCol) (transaction eacmodel.Transaction, err error) {
	if transaction, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1463467", "can't fetch the transaction", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	return
}

// List of transactions, it support pagination and search and return back count
func (p *EacTransactionServ) List(params param.Param) (transactions []eacmodel.Transaction,
	count uint64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" eac_transactions.company_id = '%v' ", params.CompanyID)
	}

	if transactions, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in transactions list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in transactions count")
	}

	return
}

// Create a transaction
func (p *EacTransactionServ) Create(transaction eacmodel.Transaction) (createdTransaction eacmodel.Transaction, err error) {

	if err = transaction.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1433872", "validation failed in creating the transaction", transaction)
		return
	}

	if createdTransaction, err = p.Repo.Create(transaction); err != nil {
		err = corerr.Tick(err, "E1479603", "transaction not created", transaction)
		return
	}

	return
}

// Save a transaction, if it is exist update it, if not create it
func (p *EacTransactionServ) Save(transaction eacmodel.Transaction) (savedTransaction eacmodel.Transaction, err error) {
	if err = transaction.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1414478", corerr.ValidationFailed, transaction)
		return
	}

	if savedTransaction, err = p.Repo.Save(transaction); err != nil {
		err = corerr.Tick(err, "E1482909", "transaction not saved")
		return
	}

	return
}

// Delete transaction, it is soft delete
func (p *EacTransactionServ) Delete(fix types.FixedCol) (transaction eacmodel.Transaction, err error) {
	if transaction, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1431984", "transaction not found for deleting")
		return
	}

	if err = p.Repo.Delete(transaction); err != nil {
		err = corerr.Tick(err, "E1484868", "transaction not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *EacTransactionServ) Excel(params param.Param) (transactions []eacmodel.Transaction, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", eacmodel.TransactionTable)

	if transactions, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1426905", "cant generate the excel list for transactions")
		return
	}

	return
}
