package service

import (
	"fmt"
	"net/http"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/action"
	"omega/internal/param"
	"omega/internal/types"
	"omega/utils/password"
)

// BasUserServ for injecting auth basrepo
type BasUserServ struct {
	Repo   basrepo.BasUserRepo
	Engine *core.Engine
}

// ProvideBasUserService for user is used in wire
func ProvideBasUserService(p basrepo.BasUserRepo) BasUserServ {
	return BasUserServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting user by it's id
func (p *BasUserServ) FindByID(id types.RowID) (user basmodel.BasUser, err error) {
	if user, err = p.Repo.FindByID(id); err != nil {
		p.Engine.CheckError(err, fmt.Sprintf("BasUser with id %v", id))
		return
	}

	return
}

// FindByUsername find user with username
func (p *BasUserServ) FindByUsername(username string) (user basmodel.BasUser, err error) {
	user, err = p.Repo.FindByUsername(username)
	p.Engine.CheckError(err, fmt.Sprintf("BasUser with username %v", username))

	return
}

// List of users, it support pagination and search and return back count
func (p *BasUserServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	p.Engine.Debug(params)

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "users list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "users count")

	return
}

func (p *BasUserServ) Create(user basmodel.BasUser,
	params param.Param) (createdBasUser basmodel.BasUser, err error) {

	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>> %p \n", p.Engine)
	var oo core.Engine
	reg := p.Engine
	oo = *p.Engine
	fmt.Printf(">>>>>>>>>>>>>>>>>>>>>>>>> %p \n", &oo)
	p.Repo.Engine = &oo

	// original := p.Engine.DB
	tx := p.Engine.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			p.Engine = reg
		}
	}()
	// p.Engine.DB = tx
	oo.DB = tx

	// if createdBasUser, err = p.CreateRollback(user, params); err != nil {
	if createdBasUser, err = p.Repo.Create(user); err != nil {
		tx.Rollback()
		p.Engine = reg
		return
	}

	// time.Sleep(30 * time.Second)
	tx.Rollback()

	// tx.Commit()
	// p.Engine.DB = original
	p.Engine = reg

	return
}

func (p *BasUserServ) CreateRollback(user basmodel.BasUser,
	params param.Param) (createdBasUser basmodel.BasUser, err error) {

	if err = user.Validate(action.Create); err != nil {
		p.Engine.CheckError(err, "Failed in validation")
		return
	}

	user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))

	if createdBasUser, err = p.Repo.Create(user); err != nil {
		// tx.Rollback()
		p.Engine.DB.Rollback()
		p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}
	// tx.Commit()
	// p.Engine.DB = original

	createdBasUser.Password = ""

	return
}

// Save user
func (p *BasUserServ) Save(user basmodel.BasUser) (createdBasUser basmodel.BasUser, err error) {

	var oldBasUser basmodel.BasUser
	oldBasUser, _ = p.FindByID(user.ID)

	if user.ID > 0 {
		if err = user.Validate(action.Update); err != nil {
			p.Engine.Debug(err)
			return
		}

		if user.Password != "" {
			user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
			p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))
		} else {
			user.Password = oldBasUser.Password
		}

	} else {
		if err = user.Validate(action.Create); err != nil {
			p.Engine.Debug(err)
			return
		}
		user.Password, err = password.Hash(user.Password, p.Engine.Envs[base.PasswordSalt])
		p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", user))
	}

	if createdBasUser, err = p.Repo.Update(user); err != nil {
		p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving user for %+v", user))
	}

	BasAccessDeleteFromCache(user.ID)

	createdBasUser.Password = ""

	return
}

// Excel is used for export excel file
func (p *BasUserServ) Excel(params param.Param) (users []basmodel.BasUser, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = "bas_users.id ASC"

	users, err = p.Repo.List(params)
	p.Engine.CheckError(err, "users excel")

	return
}

// Delete user, it is hard delete, by deleting account related to the user
func (p *BasUserServ) Delete(userID types.RowID, params param.Param) (user basmodel.BasUser, err error) {
	if user, err = p.FindByID(userID); err != nil {
		return user, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	if err = p.Repo.Delete(user); err != nil {
		p.Engine.CheckError(err, fmt.Sprintf("error in deleting user %+v", user))
	}

	return
}
