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

// BasRoleRepo for injecting engine
type BasRoleRepo struct {
	Engine *core.Engine
}

// ProvideBasRoleRepo is used in wire
func ProvideBasRoleRepo(engine *core.Engine) BasRoleRepo {
	return BasRoleRepo{Engine: engine}
}

// FindByID for role
func (p *BasRoleRepo) FindByID(id types.RowID) (role basmodel.BasRole, err error) {
	err = p.Engine.DB.First(&role, id.ToUint64()).Error
	return
}

// List of roles
func (p *BasRoleRepo) List(params param.Param) (roles []basmodel.BasRole, err error) {
	columns, err := basmodel.BasRole{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Select(columns).
		Where(search.Parse(params, basmodel.BasRole{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&roles).Error

	return
}

// Count of roles
func (p *BasRoleRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("bas_roles").
		Select(params.Select).
		Where(search.Parse(params, basmodel.BasRole{}.Pattern())).
		Count(&count).Error
	return
}

// Update BasRoleRepo
func (p *BasRoleRepo) Update(role basmodel.BasRole) (u basmodel.BasRole, err error) {
	err = p.Engine.DB.Save(&role).Error
	p.Engine.DB.Where("id = ?", role.ID).Find(&u)
	return
}

// Create BasRoleRepo
func (p *BasRoleRepo) Create(role basmodel.BasRole) (u basmodel.BasRole, err error) {
	err = p.Engine.DB.Create(&role).Scan(&u).Error
	return
}

// LastBasRole of role table
func (p *BasRoleRepo) LastBasRole(prefix types.RowID) (role basmodel.BasRole, err error) {
	err = p.Engine.DB.Unscoped().Where("id LIKE ?", fmt.Sprintf("%v%%", prefix)).
		Last(&role).Error
	if gorm.IsRecordNotFoundError(err) {
		err = nil
	}
	return
}

// Delete role
func (p *BasRoleRepo) Delete(role basmodel.BasRole) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&role).Error
	return
}
