package service

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accountstatus"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// BasAccountServ for injecting auth basrepo
type BasAccountServ struct {
	Repo   basrepo.AccountRepo
	Engine *core.Engine
}

// ProvideBasAccountService for account is used in wire
func ProvideBasAccountService(p basrepo.AccountRepo) BasAccountServ {
	return BasAccountServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting account by it's id
func (p *BasAccountServ) FindByID(fix types.FixedNode) (account basmodel.Account, err error) {
	if account, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1049049", "can't fetch the account", fix.ID, fix.CompanyID, fix.NodeID)
		return
	}

	return
}

// List of accounts, it support pagination and search and return back count
func (p *BasAccountServ) List(params param.Param) (accounts []basmodel.Account,
	count uint64, err error) {

	if accounts, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in accounts list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in accounts count")
	}

	return
}

// Create a account
func (p *BasAccountServ) Create(account basmodel.Account) (createdAccount basmodel.Account, err error) {

	if err = account.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1076780", "validation failed in creating the account", account)
		return
	}

	if createdAccount, err = p.Repo.Create(account); err != nil {
		err = corerr.Tick(err, "E1065508", "account not created", account)
		return
	}

	return
}

// Save a account, if it is exist update it, if not create it
func (p *BasAccountServ) Save(account basmodel.Account) (savedAccount basmodel.Account, err error) {
	if err = account.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1064761", corerr.ValidationFailed, account)
		return
	}

	if savedAccount, err = p.Repo.Save(account); err != nil {
		err = corerr.Tick(err, "E1084087", "account not saved")
		return
	}

	return
}

// Delete account, it is soft delete
func (p *BasAccountServ) Delete(fix types.FixedNode) (account basmodel.Account, err error) {
	if account, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1038835", "account not found for deleting")
		return
	}

	if err = p.Repo.Delete(account); err != nil {
		err = corerr.Tick(err, "E1045410", "account not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *BasAccountServ) Excel(params param.Param) (accounts []basmodel.Account, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", basmodel.AccountTable)

	if accounts, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1023076", "cant generate the excel list for accounts")
		return
	}

	return
}

// IsActive check the status of an account
func (p *BasAccountServ) IsActive(fix types.FixedNode) (bool, basmodel.Account, error) {
	var account basmodel.Account
	var err error
	if account, err = p.FindByID(fix); err != nil {
		return false, account, corerr.Tick(err, "E1059307", "account not exist", fix.ID, fix.CompanyID, fix.NodeID)
	}

	return account.Status == accountstatus.Active, account, nil
}
