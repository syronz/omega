package service

import (
	"fmt"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/enum/accountstatus"
	"omega/domain/base/enum/accounttype"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
	"omega/pkg/limberr"
	"omega/pkg/password"
)

// BasUserServ for injecting auth basrepo
type BasUserServ struct {
	Repo   basrepo.UserRepo
	Engine *core.Engine
}

// ProvideBasUserService for user is used in wire
func ProvideBasUserService(p basrepo.UserRepo) BasUserServ {
	return BasUserServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting user by it's id
func (p *BasUserServ) FindByID(fix types.FixedCol) (user basmodel.User, err error) {
	if user, err = p.Repo.FindByID(fix); err != nil {
		// err = corerr.Tick(err, "E1066324", "can't fetch the user", fix.CompanyID, fix.NodeID, fix.ID)
		err = limberr.AddCode(err, "E1066324")
		return
	}

	return
}

// FindByUsername find user with username, used for auth
func (p *BasUserServ) FindByUsername(username string) (user basmodel.User, err error) {
	if user, err = p.Repo.FindByUsername(username); err != nil {
		err = corerr.Tick(err, "E1088844", "can't fetch the user by username", username)
		return
	}

	return
}

// List of users, it support pagination and search and return back count
func (p *BasUserServ) List(params param.Param) (users []basmodel.User,
	count uint64, err error) {

	if users, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in users list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in users count")
	}

	return
}

func (p *BasUserServ) Create_Deprecated(user basmodel.User) (createdUser basmodel.User, err error) {

	if err = user.Validate(coract.Create); err != nil {
		err = corerr.TickValidate(err, "E1043810", "validatation failed in creating user", user)
		return
	}

	clonedEngine := p.Engine.Clone()
	clonedEngine.DB = clonedEngine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"users table"), "rollback recover create user")
			clonedEngine.DB.Rollback()
		}
	}()

	userRepo := basrepo.ProvideUserRepo(clonedEngine)
	accountServ := ProvideBasAccountService(basrepo.ProvideAccountRepo(clonedEngine))

	account := basmodel.Account{
		Name:   user.Name,
		Type:   accounttype.User,
		Status: accountstatus.Inactive,
	}
	account.CompanyID = user.CompanyID
	account.NodeID = user.NodeID

	var createdAccount basmodel.Account
	if createdAccount, err = accountServ.Create(account); err != nil {
		err = corerr.Tick(err, "E1067890", "error in creating account for user", user)

		clonedEngine.DB.Rollback()
		return
	}

	user.ID = createdAccount.ID
	user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
	glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	if createdUser, err = userRepo.Create(user); err != nil {
		err = corerr.Tick(err, "E1036118", "error in creating user", user)

		clonedEngine.DB.Rollback()
		return
	}

	clonedEngine.DB.Commit()
	createdUser.Password = ""

	return
}

// Create a user
func (p *BasUserServ) Create(user basmodel.User) (createdUser basmodel.User, err error) {

	if err = user.Validate(coract.Create); err != nil {
		err = corerr.TickValidate(err, "E1043810", "validatation failed in creating user", user)
		return
	}

	db := p.Engine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"users table"), "rollback recover create user")
			db.Rollback()
		}
	}()

	accountServ := ProvideBasAccountService(basrepo.ProvideAccountRepo(p.Engine))
	account := basmodel.Account{
		Name:   user.Name,
		Type:   accounttype.User,
		Status: accountstatus.Inactive,
	}
	account.CompanyID = user.CompanyID
	account.NodeID = user.NodeID

	var createdAccount basmodel.Account
	if createdAccount, err = accountServ.TxCreate(db, account); err != nil {
		err = corerr.Tick(err, "E1032795", "error in creating account for user", user)

		db.Rollback()
		return
	}

	user.ID = createdAccount.ID
	user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
	glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	if createdUser, err = p.Repo.TxCreate(db, user); err != nil {
		err = corerr.Tick(err, "E1064180", "error in creating user", user)

		db.Rollback()
		return
	}

	db.Commit()
	createdUser.Password = ""

	return

}

// Save user
func (p *BasUserServ) Save(user basmodel.User) (createdUser basmodel.User, err error) {
	var oldUser basmodel.User
	fix := types.FixedCol{
		CompanyID: user.CompanyID,
		NodeID:    user.NodeID,
		ID:        user.ID,
	}
	oldUser, _ = p.FindByID(fix)

	if err = user.Validate(coract.Update); err != nil {
		err = corerr.TickValidate(err, "E1098252", corerr.ValidationFailed, user)
		return
	}

	clonedEngine := p.Engine.Clone()
	clonedEngine.DB = clonedEngine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"users table"), "rollback recover save user")
			clonedEngine.DB.Rollback()
		}
	}()

	if user.Password != "" {
		if user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt]); err != nil {
			err = corerr.Tick(err, "E1057832", "error in saving user", user)
		}
	} else {
		user.Password = oldUser.Password
	}

	userRepo := basrepo.ProvideUserRepo(clonedEngine)
	accountServ := ProvideBasAccountService(basrepo.ProvideAccountRepo(clonedEngine))

	account := basmodel.Account{
		Name:   user.Name,
		Type:   accounttype.User,
		Status: accountstatus.Inactive,
	}
	account.ID = user.ID
	account.CompanyID = user.CompanyID
	account.NodeID = user.NodeID

	if _, err = accountServ.Save(account); err != nil {
		err = corerr.Tick(err, "E1098648", "error in saving account inside the user", user)

		clonedEngine.DB.Rollback()
		return
	}

	if createdUser, err = userRepo.Save(user); err != nil {
		err = corerr.Tick(err, "E1062983", "error in saving user", user)

		clonedEngine.DB.Rollback()
		return
	}

	BasAccessDeleteFromCache(user.ID)

	clonedEngine.DB.Commit()
	createdUser.Password = ""

	return
}

// Delete user, it is hard delete, by deleting account related to the user
func (p *BasUserServ) Delete(fix types.FixedCol) (user basmodel.User, err error) {
	if user, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1031839", "user not found for deleting")
		return
	}

	if err = p.Repo.Delete(user); err != nil {
		err = corerr.Tick(err, "E1088344", "user not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *BasUserServ) Excel(params param.Param) (users []basmodel.User, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", basmodel.UserTable)

	if users, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1064328", "cant generate the excel list")
	}

	return
}
