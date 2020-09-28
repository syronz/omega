package basrepo

import (
	// "github.com/cockroachdb/errors"

	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"reflect"
)

// UserRepo for injecting engine
type UserRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideUserRepo is used in wire
func ProvideUserRepo(engine *core.Engine) UserRepo {
	return UserRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(basmodel.User{}), basmodel.UserTable),
	}
}

// FindByID for user
func (p *UserRepo) FindByID(id types.RowID) (user basmodel.User, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).First(&user, id.ToUint64()).Error

	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		break
	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, "E1063251", corterm.ID, id, corterm.Roles)
	default:
		err = corerr.InternalServerErrorHelper(err, "E1063252")
	}

	return
}

// FindByUsername for user
func (p *UserRepo) FindByUsername(username string) (user basmodel.User, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Table("bas_users").
		Select("bas_users.*, bas_roles.resources, bas_roles.name as role").
		Where("bas_users.username = ?", username).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Scan(&user).Error

	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		break
	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, "E1021865", corterm.Username, username, corterm.Roles)
	default:
		err = corerr.InternalServerErrorHelper(err, "E1021866")
	}

	return
}

// List of users
func (p *UserRepo) List(params param.Param) (users []basmodel.User, err error) {
	columns, err := basmodel.User{}.Columns(params.Select)
	if err != nil {
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1043328").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.UserTable).Select(columns).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&users).Error

	for i := range users {
		users[i].Password = ""
	}

	return
}

// Count of users
func (p *UserRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table(basmodel.UserTable).Table("bas_users").
		Select(params.Select).
		Joins("INNER JOIN bas_roles on bas_roles.id = bas_users.role_id").
		// Where(search.Parse(params, basmodel.User{}.Pattern())).
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
