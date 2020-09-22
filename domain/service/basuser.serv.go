package service

import (
	"fmt"
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
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
	return BasUserServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting user by it's id
func (p *BasUserServ) FindByID(id types.RowID) (user basmodel.User, err error) {
	if user, err = p.Repo.FindByID(id); err != nil {
		err = corerr.Tick(err, "E1066324", "can't fetch the user")
		return
	}

	return
}

// FindByUsername find user with username
func (p *BasUserServ) FindByUsername(username string) (user basmodel.User, err error) {
	user, err = p.Repo.FindByUsername(username)
	glog.CheckInfo(err, fmt.Sprintf("User with username %v", username))

	return
}

// List of users, it support pagination and search and return back count
func (p *BasUserServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	glog.CheckError(err, "users list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	glog.CheckError(err, "users count")

	return
}

// Create a new user
func (p *BasUserServ) Create(user basmodel.User,
	params param.Param) (createdUser basmodel.User, err error) {

	if err = user.Validate(coract.Create); err != nil {
		err = limberr.Take(err, "E1043810").
			Custom(corerr.ValidationFailedErr).Build()
		glog.CheckInfo(err, corerr.ValidationFailed)
		return
	}

	clonedEngine := p.Engine.Clone()
	clonedEngine.DB = clonedEngine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v", basmodel.UserPart), "rollback recover")
			clonedEngine.DB.Rollback()
		}
	}()

	userRepo := basrepo.ProvideUserRepo(clonedEngine)

	if createdUser, err = userRepo.Create(user); err != nil {
		glog.Debug(err, createdUser)

		// if strings.Contains(strings.ToUpper(err.Error()), "FOREIGN") {
		// 	err = limberr.AddCode(err, "E1098312")
		// 	err = limberr.AddMessage(err, "database error")
		// 	err = limberr.AddType(err, "http://54323452", corerr.DuplicateHappened)
		// 	err = limberr.AddDomain(err, base.Domain)
		// 	clonedEngine.DB.Rollback()
		// 	return
		// }

		// if strings.Contains(strings.ToUpper(err.Error()), "DUPLICATE") {
		// 	clonedEngine.DB.Rollback()
		// 	return
		// }

		clonedEngine.DB.Rollback()
		return
	}

	clonedEngine.DB.Commit()

	return
}

// Save user
func (p *BasUserServ) Save(user basmodel.User, params param.Param) (createdUser basmodel.User, err error) {

	var oldUser basmodel.User
	oldUser, _ = p.FindByID(user.ID)

	if user.ID > 0 {
		if err = user.Validate(coract.Update); err != nil {
			err = limberr.Take(err, "E1098252").
				Custom(corerr.ValidationFailedErr).Build()
			glog.CheckInfo(err, corerr.ValidationFailed)
			return
		}

		if user.Password != "" {
			user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
			glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))
		} else {
			user.Password = oldUser.Password
		}

	} else {
		if err = user.Validate(coract.Create); err != nil {
			err = limberr.Take(err, "E1036447").
				Custom(corerr.ValidationFailedErr).Build()
			glog.CheckInfo(err, corerr.ValidationFailed)
			return
		}
		user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
		glog.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))
	}

	if createdUser, err = p.Repo.Update(user); err != nil {
		glog.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}

	BasAccessDeleteFromCache(user.ID)

	createdUser.Password = ""

	return
}

// Excel is used for export excel file
func (p *BasUserServ) Excel(params param.Param) (users []basmodel.User, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = "bas_users.id ASC"

	users, err = p.Repo.List(params)
	glog.CheckError(err, "users excel")

	return
}

// Delete user, it is hard delete, by deleting account related to the user
func (p *BasUserServ) Delete(userID types.RowID) (user basmodel.User, err error) {
	if user, err = p.FindByID(userID); err != nil {
		return user, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	if err = p.Repo.Delete(user); err != nil {
		glog.CheckError(err, fmt.Sprintf("error in deleting user %+v", user))
	}

	return
}
