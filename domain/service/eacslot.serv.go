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

// EacSlotServ for injecting auth eacrepo
type EacSlotServ struct {
	Repo         eacrepo.SlotRepo
	Engine       *core.Engine
	CurrencyServ EacCurrencyServ
	AccountServ  BasAccountServ
}

// ProvideEacSlotService for slot is used in wire
func ProvideEacSlotService(p eacrepo.SlotRepo, currencyServ EacCurrencyServ,
	accountServ BasAccountServ) EacSlotServ {
	return EacSlotServ{
		Repo:         p,
		Engine:       p.Engine,
		CurrencyServ: currencyServ,
		AccountServ:  accountServ,
	}
}

// FindByID for getting slot by it's id
func (p *EacSlotServ) FindByID(fix types.FixedCol) (slot eacmodel.Slot, err error) {
	if slot, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1410775", "can't fetch the slot", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	return
}

// List of slots, it support pagination and search and return back count
func (p *EacSlotServ) List(params param.Param) (slots []eacmodel.Slot,
	count uint64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" eac_slots.company_id = '%v' ", params.CompanyID)
	}

	if slots, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in slots list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in slots count")
	}

	return
}

// TransactionSlot is used inside Transaction.FindByID
func (p *EacSlotServ) TransactionSlot(transactionID types.RowID,
	companyID, nodeID uint64) (slots []eacmodel.Slot,
	err error) {

	params := param.New()
	params.PreCondition = fmt.Sprintf(" eac_slots.company_id = '%v' AND eac_slots.transaction_id = '%v' AND eac_slots.node_id = '%v' ",
		companyID, transactionID, nodeID)

	if slots, err = p.Repo.List(params); err != nil {
		return
	}

	return
}

// Create a slot
func (p *EacSlotServ) Create(slot eacmodel.Slot) (createdSlot eacmodel.Slot, err error) {
	if err = slot.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1437242", "validation failed in creating the slot", slot)
		return
	}

	var lastSlot eacmodel.Slot
	if lastSlot, err = p.Repo.LastSlot(slot); err != nil {
		err = corerr.Tick(err, "E1445069", "last slot not found", slot)
		return
	}

	adjust := slot.Debit - slot.Credit
	slot.Balance = lastSlot.Balance + adjust

	if createdSlot, err = p.Repo.Create(slot); err != nil {
		err = corerr.Tick(err, "E1434523", "slot not created", slot)
		return
	}

	if err = p.Repo.RegulateBalances(slot, adjust); err != nil {
		err = corerr.Tick(err, "E1466626", "regulate balances faced error in create", slot, adjust)
		return
	}

	return
}

// Reset will remove affect of transaction on journal, similar to delete but don't delete the
// records
func (p *EacSlotServ) Reset(slot eacmodel.Slot) (err error) {
	adjust := slot.Credit - slot.Debit
	slot.Debit = 0
	slot.Credit = 0
	slot.Balance = 0
	if slot, err = p.Repo.Save(slot); err != nil {
		err = corerr.Tick(err, "E1445231", "error in resetting the slot", slot)
		return
	}

	if err = p.Repo.RegulateBalancesSave(slot, adjust); err != nil {
		err = corerr.Tick(err, "E1491272", "regulate balances faced error in reset", slot, adjust)
		return
	}

	return
}

// Save a slot, if it is exist update it, if not create it
func (p *EacSlotServ) Save(slot eacmodel.Slot) (savedSlot eacmodel.Slot, err error) {
	if err = slot.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1445816", corerr.ValidationFailed, slot)
		return
	}

	fix := types.FixedCol{
		CompanyID: slot.CompanyID,
		NodeID:    slot.NodeID,
		ID:        slot.ID,
	}

	var oldSlot eacmodel.Slot

	if oldSlot, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1482909", "slot not found", slot)
		return
	}

	p.Reset(oldSlot)

	var lastSlot eacmodel.Slot
	if lastSlot, err = p.Repo.LastSlotWithID(slot); err != nil {
		err = corerr.Tick(err, "E1433617", "last slot not found in save transaction", slot)
		return
	}

	adjust := slot.Debit - slot.Credit
	slot.Balance = lastSlot.Balance + adjust

	if slot, err = p.Repo.Save(slot); err != nil {
		err = corerr.Tick(err, "E1475746", "error in saving the slot", slot)
		return
	}

	if err = p.Repo.RegulateBalancesSave(slot, adjust); err != nil {
		err = corerr.Tick(err, "E1454858", "regulate balances faced error in save", slot, adjust)
		return
	}

	// var adjust float64

	// if oldSlot.AccountID != slot.AccountID {
	// 	adjust = oldSlot.Credit - oldSlot.Debit
	// 	if err = p.Repo.RegulateBalancesSave(oldSlot, adjust); err != nil {
	// 		err = corerr.Tick(err, "E1483419", "regulate balances faced error in save", slot, adjust)
	// 		return
	// 	}

	// 	var lastSlot eacmodel.Slot
	// 	if lastSlot, err = p.Repo.LastSlot(slot); err != nil {
	// 		err = corerr.Tick(err, "", "last slot not found for save", slot)
	// 		return
	// 	}
	// 	adjust = slot.Debit - slot.Credit
	// 	slot.Balance = lastSlot.Balance + adjust
	// } else {
	// 	slot.Balance = oldSlot.Balance - oldSlot.Debit + oldSlot.Credit - slot.Credit + slot.Debit
	// 	adjust = slot.Debit - slot.Credit - oldSlot.Balance
	// }

	// if savedSlot, err = p.Repo.Save(slot); err != nil {
	// 	err = corerr.Tick(err, "E1434918", "slot not saved")
	// 	return
	// }

	// if err = p.Repo.RegulateBalancesSave(slot, adjust); err != nil {
	// 	err = corerr.Tick(err, "E1421914", "regulate balances faced error in save", slot, adjust)
	// 	return
	// }

	return
}

// Delete slot, it is soft delete
func (p *EacSlotServ) Delete(fix types.FixedCol) (slot eacmodel.Slot, err error) {
	if slot, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1486984", "slot not found for deleting")
		return
	}

	if err = p.Repo.Delete(slot); err != nil {
		err = corerr.Tick(err, "E1482264", "slot not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *EacSlotServ) Excel(params param.Param) (slots []eacmodel.Slot, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", eacmodel.SlotTable)

	if slots, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1485891", "cant generate the excel list for slots")
		return
	}

	return
}
