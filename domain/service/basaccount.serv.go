package service

import (
	"encoding/json"
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accountstatus"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"

	"gorm.io/gorm"
)

// BasAccountServ for injecting auth basrepo
type BasAccountServ struct {
	Repo      basrepo.AccountRepo
	Engine    *core.Engine
	PhoneServ BasPhoneServ
}

// ProvideBasAccountService for account is used in wire
func ProvideBasAccountService(p basrepo.AccountRepo, phoneServ BasPhoneServ) BasAccountServ {
	return BasAccountServ{
		Repo:      p,
		Engine:    p.Engine,
		PhoneServ: phoneServ,
	}
}

// FindByID for getting account by it's id
func (p *BasAccountServ) FindByID(fix types.FixedNode) (account basmodel.Account, err error) {
	if account, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1049049", "can't fetch the account", fix.ID, fix.CompanyID, fix.NodeID)
		return
	}

	if account.Phones, err = p.PhoneServ.AccountsPhones(fix); err != nil {
		err = corerr.Tick(err, "E1017084", "can't fetch the account's phones", fix.ID, fix.CompanyID, fix.NodeID)
		return

	}

	return
}

// List of accounts, it support pagination and search and return back count
func (p *BasAccountServ) List(params param.Param) (accounts []basmodel.Account,
	count int64, err error) {

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
	db := p.Engine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"users table"), "rollback recover create user")
			db.Rollback()
		}
	}()

	if createdAccount, err = p.TxCreate(p.Repo.Engine.DB, account); err != nil {
		err = corerr.Tick(err, "E1014394", "error in creating account for user", createdAccount)

		db.Rollback()
		return
	}

	db.Commit()

	return
}

// TxCreate is used for creating an account in case of transaction activated
func (p *BasAccountServ) TxCreate(db *gorm.DB, account basmodel.Account) (createdAccount basmodel.Account, err error) {
	if err = account.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1076780", "validation failed in creating the account", account)
		return
	}

	if createdAccount, err = p.Repo.TxCreate(db, account); err != nil {
		err = corerr.Tick(err, "E1065508", "account not created", account)
		return
	}

	phoneServ := ProvideBasPhoneService(basrepo.ProvidePhoneRepo(p.Engine))

	for _, phone := range account.Phones {
		phone.CompanyID = createdAccount.CompanyID
		phone.NodeID = createdAccount.NodeID
		phone.AccountID = createdAccount.ID
		if _, err = phoneServ.TxCreate(db, phone); err != nil {
			err = corerr.Tick(err, "E1040913", "error in creating phone for account", phone)

			return
		}
	}

	return
}

// Save a account, if it is exist update it, if not create it
func (p *BasAccountServ) Save(account basmodel.Account) (savedAccount basmodel.Account, err error) {
	return p.TxSave(p.Engine.DB, account)
}

// TxSave a account, if it is exist update it, if not create it
func (p *BasAccountServ) TxSave(db *gorm.DB, account basmodel.Account) (savedAccount basmodel.Account, err error) {
	if err = account.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1064761", corerr.ValidationFailed, account)
		return
	}

	if savedAccount, err = p.Repo.TxSave(db, account); err != nil {
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
	params.Limit = p.Engine.Envs.ToInt(core.ExcelMaxRows)
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

func makeTreeChartOfAccounts(accounts []basmodel.Account) {
	arr := make([]basmodel.Tree, len(accounts))

	for i, v := range accounts {
		arr[i].ID = v.ID
		arr[i].CompanyID = v.CompanyID
		arr[i].NodeID = v.NodeID
		arr[i].ParentID = v.ParentID
		arr[i].Code = v.Code
		arr[i].Name = v.Name
		arr[i].Type = v.Type
	}

	pMap := make(map[types.RowID]*basmodel.Tree, 1)

	var root basmodel.Tree
	pMap[0] = &root

	exceed := basmodel.Tree{
		Name: "exceed",
	}

	for i, v := range arr {
		pMap[v.ID] = &arr[i]

		pID := parseParent(v.ParentID)

		pMap[pID].Counter++
		if pMap[pID].Counter < consts.MaxChildrenForChartOfAccounts {
			pMap[pID].Children = append(pMap[pID].Children, &arr[i])
		} else {
			if pMap[pID].Counter == consts.MaxChildrenForChartOfAccounts {
				exceed.ParentID = v.ParentID
				pMap[pID].Children = append(pMap[pID].Children, &exceed)
			}
		}

	}

	b, err := json.MarshalIndent(root, "", "    ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

func parseParent(pID *types.RowID) types.RowID {
	if pID == nil {
		return 0
	}
	return *pID
}
