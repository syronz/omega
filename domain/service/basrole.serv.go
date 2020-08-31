package service

import (
	"fmt"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/term"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/glog"
	"strings"

	"github.com/jinzhu/gorm"
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
func (p *BasRoleServ) FindByID(params param.Param, id types.RowID) (role basmodel.Role, err error) {
	role, err = p.Repo.FindByID(id)

	if gorm.IsRecordNotFoundError(err) {
		err = corerr.New("E1032412", params, base.Domain, err, id).
			NotFound(basmodel.RolesPart, "id", id, "roles/"+id.ToString())
		return
	}

	if err != nil {
		err = corerr.New("E1032423", params, base.Domain, err, id).
			InternalServer("roles/" + id.ToString())
		return
	}
	// glog.CheckError(err, fmt.Sprintf("Role with id %v", id))

	return
}

// List of roles, it support pagination and search and return back count
func (p *BasRoleServ) List(params param.Param) (roles []basmodel.Role,
	count uint64, err error) {

	// data = make(map[string]interface{})

	if roles, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "roles list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "roles count")
	}

	return
}

// Create a role
func (p *BasRoleServ) Create(role basmodel.Role, params param.Param) (createdRole basmodel.Role, err error) {

	if err = role.Validate(coract.Save, params); err != nil {
		glog.CheckError(err, term.Validation_failed)
		return
	}

	if createdRole, err = p.Repo.Create(role); err != nil {
		if strings.Contains(strings.ToUpper(err.Error()), "DUPLICATE") {
			err = corerr.New("E1074134", params, base.Domain, err, role).
				FieldError("/roles", corerr.Duplication_happened).
				Add("name", corerr.This_V_already_exist, dict.R("name"))
			return
		}
		err = corerr.New("E10522393", params, base.Domain, err, role.Name, role.Resources, role.Description).
			InternalServer("/roles")
		return
	}

	return
}

// Save a role, if it is exist update it, if not create it
func (p *BasRoleServ) Save(role basmodel.Role) (savedRole basmodel.Role, err error) {
	var params param.Param

	if err = role.Validate(coract.Save, params); err != nil {
		glog.CheckError(err, "validation failed")
		return
	}

	savedRole, err = p.Repo.Update(role)
	glog.CheckInfo(err, fmt.Sprintf("Failed in updating role for %+v", role))
	if err == nil {
		BasAccessResetFullCache()
	}

	return
}

// Delete role, it is soft delete
func (p *BasRoleServ) Delete(roleID types.RowID, params param.Param) (role basmodel.Role, err error) {

	if role, err = p.FindByID(params, roleID); err != nil {
		return
	}

	err = p.Repo.Delete(role)
	if err == nil {
		BasAccessResetFullCache()
	}
	return
}

// Excel is used for export excel file
func (p *BasRoleServ) Excel(params param.Param) (roles []basmodel.Role, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = "bas_roles.id ASC"

	roles, err = p.Repo.List(params)
	glog.CheckError(err, "roles excel")

	return
}
