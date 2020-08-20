package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/search"
	"omega/internal/types"
)

// BasUserRepo for injecting engine
type BasUserRepo struct {
	Engine *core.Engine
}

// ProvideBasUserRepo is used in wire
func ProvideBasUserRepo(engine *core.Engine) BasUserRepo {
	return BasUserRepo{Engine: engine}
}

// FindByID for user
func (p *BasUserRepo) FindByID(id types.RowID) (user basmodel.BasUser, err error) {
	err = p.Engine.DB.First(&user, id.ToUint64()).Error
	return
}

// FindByUsername for user
func (p *BasUserRepo) FindByUsername(username string) (user basmodel.BasUser, err error) {
	err = p.Engine.DB.Table("bas_users").
		Select("bas_users.*, bas_roles.resources").
		Where("bas_users.username = ?", username).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Scan(&user).Error
	return
}

// List of users
func (p *BasUserRepo) List(params param.Param) (users []basmodel.BasUser, err error) {
	columns, err := basmodel.BasUser{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table("bas_users").Select(columns).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Where(search.Parse(params, basmodel.BasUser{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&users).Error

	return
}

// Count of users
func (p *BasUserRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("bas_users").
		Select(params.Select).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Where(search.Parse(params, basmodel.BasUser{}.Pattern())).
		Count(&count).Error
	return
}

// Update BasUserRepo
func (p *BasUserRepo) Update(user basmodel.BasUser) (u basmodel.BasUser, err error) {
	err = p.Engine.DB.Save(&user).Error
	p.Engine.DB.Where("id = ?", user.ID).Find(&u)
	return
}

// Create BasUserRepo
func (p *BasUserRepo) Create(user basmodel.BasUser) (u basmodel.BasUser, err error) {
	err = p.Engine.DB.Create(&user).Scan(&u).Error
	return
}

// Delete user
func (p *BasUserRepo) Delete(user basmodel.BasUser) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&user).Error
	return
}
