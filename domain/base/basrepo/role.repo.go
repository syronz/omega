package basrepo

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/param"
	"omega/internal/search"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/limberr"

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

	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		break
	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, "E1072991", corterm.ID, id, corterm.Roles)
	default:
		err = corerr.InternalServerErrorHelper(err, "E1072992")
	}
	return
}

// List of roles
func (p *RoleRepo) List(params param.Param) (roles []basmodel.Role, err error) {
	columns, err := basmodel.Role{}.Columns(params.Select, params)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table(basmodel.RoleTable).Select(columns).
		Where(search.Parse(params, basmodel.Role{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&roles).Error

	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		break
	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, "E1032861")
	default:
		err = corerr.InternalServerErrorHelper(err, "E1022879")
	}

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
	if err != nil {
		switch corerr.ClearDbErr(err) {
		case corerr.ForeignErr:
			err = limberr.Take(err, "E1054817").
				Message(corerr.SomeVRelatedToThisVSoItIsNotCreated, dict.R(corterm.Users), dict.R(corterm.Role)).
				Custom(corerr.ForeignErr).Build()
		case corerr.DuplicateErr:
			err = limberr.Take(err, "E1054818").
				Message(corerr.VWithValueVAlreadyExist, dict.R(corterm.Role), role.Name).
				Custom(corerr.DuplicateErr).Build()
			err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, role.Name)
		default:
			err = limberr.Take(err, "E1054819").
				Message(corerr.InternalServerError).
				Custom(corerr.InternalServerErr).Build()
		}
	}
	p.Engine.DB.Table(basmodel.RoleTable).Where("id = ?", role.ID).Find(&u)
	return
}

// Create RoleRepo
func (p *RoleRepo) Create(role basmodel.Role) (u basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Create(&role).Scan(&u).Error
	if err != nil {
		switch corerr.ClearDbErr(err) {
		case corerr.ForeignErr:
			err = limberr.Take(err, "E1053287").
				Message(corerr.SomeVRelatedToThisVSoItIsNotCreated, dict.R(corterm.Users), dict.R(corterm.Role)).
				Custom(corerr.ForeignErr).Build()
		case corerr.DuplicateErr:
			err = limberr.Take(err, "E1053288").
				Message(corerr.VWithValueVAlreadyExist, dict.R(corterm.Role), role.Name).
				Custom(corerr.DuplicateErr).Build()
			err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, role.Name)
		default:
			err = limberr.Take(err, "E1053289").
				Message(corerr.InternalServerError).
				Custom(corerr.InternalServerErr).Build()
		}
	}
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
	if err = p.Engine.DB.Table(basmodel.RoleTable).Unscoped().Delete(&role).Error; err != nil {
		switch corerr.ClearDbErr(err) {
		case corerr.ForeignErr:
			err = limberr.Take(err, "E1067392").
				Message(corerr.SomeVRelatedToThisVSoItIsNotDeleted, dict.R(corterm.Users), dict.R(corterm.Role)).
				Custom(corerr.ForeignErr).Build()
		default:
			err = limberr.Take(err, "E1067393").
				Message(corerr.InternalServerError).
				Custom(corerr.InternalServerErr).Build()
		}
	}
	return
}
