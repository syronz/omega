package service

import (
	"fmt"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacrepo"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
	"time"
)

// EacTransactionServ for injecting auth eacrepo
type EacTransactionServ struct {
	Repo     eacrepo.TransactionRepo
	Engine   *core.Engine
	SlotServ EacSlotServ
}

// ProvideEacTransactionService for transaction is used in wire
func ProvideEacTransactionService(p eacrepo.TransactionRepo, slotServ EacSlotServ) EacTransactionServ {
	return EacTransactionServ{
		Repo:     p,
		Engine:   p.Engine,
		SlotServ: slotServ,
	}
}

// FindByID for getting transaction by it's id
func (p *EacTransactionServ) FindByID(fix types.FixedCol) (transaction eacmodel.Transaction, err error) {
	if transaction, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1463467", "can't fetch the transaction", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	if transaction.Slots, err = p.SlotServ.TransactionSlot(transaction.ID,
		transaction.CompanyID, transaction.NodeID); err != nil {
		err = corerr.Tick(err, "E1464213", "can't fetch the transactions slots",
			fix.CompanyID, fix.NodeID, transaction.ID)
		return
	}

	return
}

// List of transactions, it support pagination and search and return back count
func (p *EacTransactionServ) List(params param.Param) (transactions []eacmodel.Transaction,
	count int64, err error) {

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

// Transfer activate create for special transfering
func (p *EacTransactionServ) Transfer(transaction eacmodel.Transaction) (createdTransaction eacmodel.Transaction, err error) {
	slots := []eacmodel.Slot{
		{
			AccountID:  transaction.Pioneer,
			Credit:     transaction.Amount,
			CurrencyID: transaction.CurrencyID,
			PostDate:   transaction.PostDate,
		},
		{
			AccountID:  transaction.Follower,
			Debit:      transaction.Amount,
			CurrencyID: transaction.CurrencyID,
			PostDate:   transaction.PostDate,
		},
	}

	slots[0].CompanyID = transaction.CompanyID
	slots[1].CompanyID = transaction.CompanyID
	slots[0].NodeID = transaction.NodeID
	slots[1].NodeID = transaction.NodeID

	createdTransaction, err = p.Create(transaction, slots)

	return
}

// Create a transaction
func (p *EacTransactionServ) Create(transaction eacmodel.Transaction,
	slots []eacmodel.Slot) (createdTransaction eacmodel.Transaction, err error) {

	if err = transaction.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1433872", "validation failed in creating the transaction", transaction)
		return
	}

	clonedEngine := p.Engine.Clone()
	clonedEngine.DB = clonedEngine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"eac_transactions table"), "rollback recover create transaction")
			clonedEngine.DB.Rollback()
		}
	}()

	transactionRepo := eacrepo.ProvideTransactionRepo(clonedEngine)
	slotServ := ProvideEacSlotService(eacrepo.ProvideSlotRepo(clonedEngine),
		p.SlotServ.CurrencyServ, p.SlotServ.AccountServ)

	now := time.Now()
	transaction.Hash = now.Format(consts.HashTimeLayout)

	if createdTransaction, err = transactionRepo.Create(transaction); err != nil {
		err = corerr.Tick(err, "E1479603", "transaction not created", transaction)

		clonedEngine.DB.Rollback()
		return
	}

	for _, v := range slots {
		v.TransactionID = createdTransaction.ID
		if _, err = slotServ.Create(v); err != nil {
			err = corerr.Tick(err, "E1420630", "slot not saved in transaction creation", v)
			clonedEngine.DB.Rollback()
			return
		}
	}

	clonedEngine.DB.Commit()

	return
}

// EditTransfer activate create for special transfering
func (p *EacTransactionServ) EditTransfer(tr eacmodel.Transaction) (updatedTr eacmodel.Transaction, err error) {
	slots := []eacmodel.Slot{
		{
			AccountID:  tr.Pioneer,
			Credit:     tr.Amount,
			CurrencyID: tr.CurrencyID,
			PostDate:   tr.PostDate,
		},
		{
			AccountID:  tr.Follower,
			Debit:      tr.Amount,
			CurrencyID: tr.CurrencyID,
			PostDate:   tr.PostDate,
		},
	}

	fix := types.FixedCol{
		CompanyID: tr.CompanyID,
		NodeID:    tr.NodeID,
		ID:        tr.ID,
	}

	var oldTr eacmodel.Transaction
	if oldTr, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1484155", "edit transfer can't find the transaction", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	slots[0].CompanyID = tr.CompanyID
	slots[0].NodeID = tr.NodeID
	slots[0].ID = oldTr.Slots[0].ID

	slots[1].CompanyID = tr.CompanyID
	slots[1].NodeID = tr.NodeID
	slots[1].ID = oldTr.Slots[1].ID

	updatedTr, err = p.Update(tr, slots)

	return
}

// Update is used when a transaction has been changed
func (p *EacTransactionServ) Update(tr eacmodel.Transaction,
	slots []eacmodel.Slot) (updatedTr eacmodel.Transaction, err error) {

	if err = tr.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1469927", "validation failed in updating the transaction", tr)
		return
	}

	clonedEngine := p.Engine.Clone()
	clonedEngine.DB = clonedEngine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"eac_transactions table"), "rollback recover update transaction")
			clonedEngine.DB.Rollback()
		}
	}()

	transactionRepo := eacrepo.ProvideTransactionRepo(clonedEngine)
	slotServ := ProvideEacSlotService(eacrepo.ProvideSlotRepo(clonedEngine),
		p.SlotServ.CurrencyServ, p.SlotServ.AccountServ)

	now := time.Now()
	tr.Hash = now.Format(consts.HashTimeLayout)

	if updatedTr, err = transactionRepo.Save(tr); err != nil {
		err = corerr.Tick(err, "E1479603", "transaction not updated", tr)

		clonedEngine.DB.Rollback()
		return
	}

	for _, v := range slots {
		v.TransactionID = updatedTr.ID
		if _, err = slotServ.Save(v); err != nil {
			err = corerr.Tick(err, "E1420630", "slot not saved in updating the transaction", v)
			clonedEngine.DB.Rollback()
			return
		}
	}

	clonedEngine.DB.Commit()

	return

}

// Save a transaction, if it is exist update it, if not create it
func (p *EacTransactionServ) Save(transaction eacmodel.Transaction) (savedTransaction eacmodel.Transaction, err error) {
	if err = transaction.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1414478", corerr.ValidationFailed, transaction)
		return
	}

	clonedEngine := p.Engine.Clone()
	clonedEngine.DB = clonedEngine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"eac_transactions table"), "rollback recover update transaction")
			clonedEngine.DB.Rollback()
		}
	}()

	transactionRepo := eacrepo.ProvideTransactionRepo(clonedEngine)
	slotServ := ProvideEacSlotService(eacrepo.ProvideSlotRepo(clonedEngine),
		p.SlotServ.CurrencyServ, p.SlotServ.AccountServ)

	if savedTransaction, err = transactionRepo.Save(transaction); err != nil {
		err = corerr.Tick(err, "E1482909", "transaction not saved", transaction)

		clonedEngine.DB.Rollback()
		return
	}

	for _, v := range transaction.Slots {
		v.PostDate = transaction.PostDate
		v.TransactionID = transaction.ID
		if _, err = slotServ.Save(v); err != nil {
			err = corerr.Tick(err, "E1485649", "slot not saved in transaction edit", v)
			clonedEngine.DB.Rollback()
			return
		}
	}

	clonedEngine.DB.Commit()

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
	params.Limit = p.Engine.Envs.ToInt(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", eacmodel.TransactionTable)

	if transactions, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1426905", "cant generate the excel list for transactions")
		return
	}

	return
}
