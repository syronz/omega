package service

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// BasRoleServ for injecting auth basrepo
type BasRoleServ struct {
	Repo   basrepo.RoleRepo
	Engine *core.Engine
}

// ProvideBasRoleService for role is used in wire
func ProvideBasRoleService(p basrepo.RoleRepo) BasRoleServ {
	return BasRoleServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting role by it's id
func (p *BasRoleServ) FindByID(id types.RowID) (role basmodel.Role, err error) {
	if role, err = p.Repo.FindByID(id); err != nil {
		err = corerr.Tick(err, "E1043183", "can't fetch the role", id)
		return
	}

	return
}

// List of roles, it support pagination and search and return back count
func (p *BasRoleServ) List(params param.Param) (roles []basmodel.Role,
	count uint64, err error) {

	if roles, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in roles list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in roles count")
	}

	return
}

// Create a role
func (p *BasRoleServ) Create(role basmodel.Role) (createdRole basmodel.Role, err error) {

	if err = role.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1098554", "validation failed in creating the role", role)
		return
	}

	if createdRole, err = p.Repo.Create(role); err != nil {
		err = corerr.Tick(err, "E1042894", "role not created", role)
		return
	}

	return
}

// Save a role, if it is exist update it, if not create it
func (p *BasRoleServ) Save(role basmodel.Role) (savedRole basmodel.Role, err error) {
	if err = role.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1037119", corerr.ValidationFailed, role)
		return
	}

	if savedRole, err = p.Repo.Save(role); err != nil {
		err = corerr.Tick(err, "E1078742", "role not saved")
		return
	}

	BasAccessResetFullCache()
	return
}

// Delete role, it is soft delete
func (p *BasRoleServ) Delete(roleID types.RowID) (role basmodel.Role, err error) {
	if role, err = p.FindByID(roleID); err != nil {
		err = corerr.Tick(err, "E1052861", "role not found for deleting")
		return
	}

	if err = p.Repo.Delete(role); err != nil {
		err = corerr.Tick(err, "E1017987", "role not deleted")
		return
	}

	BasAccessResetFullCache()
	return
}

// Excel is used for export excel file
func (p *BasRoleServ) Excel(params param.Param) (roles []basmodel.Role, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", basmodel.RoleTable)

	if roles, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1067385", "cant generate the excel list for roles")
		return
	}

	return
}
