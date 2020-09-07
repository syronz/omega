package basrepo

import (
	// "github.com/cockroachdb/errors"

	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/search"
	"omega/internal/types"
)

// UserRepo for injecting engine
type UserRepo struct {
	Engine *core.Engine
}

// ProvideUserRepo is used in wire
func ProvideUserRepo(engine *core.Engine) UserRepo {
	return UserRepo{Engine: engine}
}

// FindByID for user
func (p *UserRepo) FindByID(id types.RowID) (user basmodel.User, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).First(&user, id.ToUint64()).Error
	return
}

// FindByUsername for user
func (p *UserRepo) FindByUsername(username string) (user basmodel.User, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Table("bas_users").
		Select("bas_users.*, bas_roles.resources").
		Where("bas_users.username = ?", username).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Scan(&user).Error
	return
}

// List of users
func (p *UserRepo) List(params param.Param) (users []basmodel.User, err error) {
	columns, err := basmodel.User{}.Columns(params.Select, params)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table(basmodel.UserTable).Select(columns).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Where(search.Parse(params, basmodel.User{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&users).Error

	return
}

// Count of users
func (p *UserRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Table("bas_users").
		Select(params.Select).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Where(search.Parse(params, basmodel.User{}.Pattern())).
		Count(&count).Error
	return
}

// Update UserRepo
func (p *UserRepo) Update(user basmodel.User) (u basmodel.User, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Save(&user).Error
	p.Engine.DB.Table(basmodel.UserTable).Where("id = ?", user.ID).Find(&u)
	return
}

// Create UserRepo
func (p *UserRepo) Create(user basmodel.User) (u basmodel.User, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Create(&user).Scan(&u).Error
	return
}

// Delete user
func (p *UserRepo) Delete(user basmodel.User) (err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Unscoped().Delete(&user).Error
	return
}
