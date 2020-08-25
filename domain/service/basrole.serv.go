package service

import (
	"fmt"
	"net/http"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/action"
	"omega/internal/param"
	"omega/internal/term"
	"omega/internal/types"
)

// BasRoleServ for injecting auth basrepo
type BasRoleServ struct {
	Repo   basrepo.RoleRepo
	Engine *core.Engine
}

// ProvideBasRoleService for role is used in wire
func ProvideBasRoleService(p basrepo.RoleRepo) BasRoleServ {
	return BasRoleServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting role by it's id
func (p *BasRoleServ) FindByID(id types.RowID) (role basmodel.Role, err error) {
	role, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Role with id %v", id))

	return
}

// List of roles, it support pagination and search and return back count
func (p *BasRoleServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	data["list"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "roles list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "roles count")

	return
}

// Create a role
func (p *BasRoleServ) Create(role basmodel.Role, params param.Param) (createdRole basmodel.Role, err error) {

	if err = role.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, term.Validation_failed)
		return
	}

	createdRole, err = p.Repo.Create(role)

	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in creating role for %+v", role))

	return
}

// Save a role, if it is exist update it, if not create it
func (p *BasRoleServ) Save(role basmodel.Role) (savedRole basmodel.Role, err error) {

	if err = role.Validate(action.Save); err != nil {
		p.Engine.CheckError(err, "validation failed")
		return
	}

	savedRole, err = p.Repo.Update(role)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in updating role for %+v", role))
	if err == nil {
		BasAccessResetFullCache()
	}

	return
}

// Delete role, it is soft delete
func (p *BasRoleServ) Delete(roleID types.RowID, params param.Param) (role basmodel.Role, err error) {

	if role, err = p.FindByID(roleID); err != nil {
		return
	}

	err = p.Repo.Delete(role)
	if err == nil {
		BasAccessResetFullCache()
	}
	return
}

// HardDelete will delete the role permanently
func (p *BasRoleServ) HardDelete(roleID types.RowID) error {
	role, err := p.FindByID(roleID)
	if err != nil {
		return core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	BasAccessResetFullCache()

	return p.Repo.Delete(role)
}

// Excel is used for export excel file
func (p *BasRoleServ) Excel(params param.Param) (roles []basmodel.Role, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = "bas_roles.id ASC"

	roles, err = p.Repo.List(params)
	p.Engine.CheckError(err, "roles excel")

	return
}
