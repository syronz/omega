package basrepo

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/search"
	"omega/internal/types"

	"github.com/jinzhu/gorm"
)

// RoleRepo for injecting engine
type RoleRepo struct {
	Engine *core.Engine
}

// ProvideRoleRepo is used in wire
func ProvideRoleRepo(engine *core.Engine) RoleRepo {
	return RoleRepo{Engine: engine}
}

// FindByID for role
func (p *RoleRepo) FindByID(id types.RowID) (role basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).First(&role, id.ToUint64()).Error
	return
}

// List of roles
func (p *RoleRepo) List(params param.Param) (roles []basmodel.Role, err error) {
	columns, err := basmodel.Role{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table(basmodel.RoleTable).Select(columns).
		Where(search.Parse(params, basmodel.Role{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&roles).Error

	return
}

// Count of roles
func (p *RoleRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).
		Select(params.Select).
		Where(search.Parse(params, basmodel.Role{}.Pattern())).
		Count(&count).Error
	return
}

// Update RoleRepo
func (p *RoleRepo) Update(role basmodel.Role) (u basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Save(&role).Error
	p.Engine.DB.Where("id = ?", role.ID).Find(&u)
	return
}

// Create RoleRepo
func (p *RoleRepo) Create(role basmodel.Role) (u basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Create(&role).Scan(&u).Error
	return
}

// LastRole of role table
func (p *RoleRepo) LastRole(prefix types.RowID) (role basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Unscoped().Where("id LIKE ?", fmt.Sprintf("%v%%", prefix)).
		Last(&role).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

// Delete role
func (p *RoleRepo) Delete(role basmodel.Role) (err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Unscoped().Delete(&role).Error
	return
}
