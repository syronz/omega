package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/glog"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"reflect"
)

// RoleRepo for injecting engine
type RoleRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideRoleRepo is used in wire
func ProvideRoleRepo(engine *core.Engine) RoleRepo {
	return RoleRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(basmodel.Role{}), basmodel.RoleTable),
	}
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
	cols, err := validator.CheckColumns(p.Cols, params.Select)
	if err != nil {
		glog.Debug(err)
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1032278").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.RoleTable).Select(cols).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&roles).Error

	// switch corerr.ClearDbErr(err) {
	// case corerr.Nil:
	// 	break
	// case corerr.ValidationFailedErr:
	// 	err = corerr.ValidationFailedHelper(err, "E1032861")
	// default:
	// 	err = corerr.InternalServerErrorHelper(err, "E1022879")
	// }

	err = p.dbError(err, "E1032861", 

	return
}

// Count of roles
func (p *RoleRepo) Count(params param.Param) (count uint64, err error) {

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1032288").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.RoleTable).
		Where(whereStr).
		Count(&count).Error
	return
}

// Update RoleRepo
func (p *RoleRepo) Update(role basmodel.Role) (u basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Save(&role).Error

	if err != nil {
		err = p.dbError(err, "E1054817", role, corterm.Updated)
	}

	p.Engine.DB.Table(basmodel.RoleTable).Where("id = ?", role.ID).Find(&u)
	return
}

// Create RoleRepo
func (p *RoleRepo) Create(role basmodel.Role) (u basmodel.Role, err error) {
	err = p.Engine.DB.Table(basmodel.RoleTable).Create(&role).Scan(&u).Error
	if err != nil {
		err = p.dbError(err, "E1053287", role, corterm.Created)
	}
	return
}

// Delete role
func (p *RoleRepo) Delete(role basmodel.Role) (err error) {
	if err = p.Engine.DB.Table(basmodel.RoleTable).Unscoped().Delete(&role).Error; err != nil {
		err = p.dbError(err, "E1067392", role, corterm.Deleted)
	}
	return
}

// dbError is an internal method for create proper database error
func (p *RoleRepo) dbError(err error, code string, role basmodel.Role, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil
		break
	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(corterm.Users),
				dict.R(corterm.Role), dict.R(action)).
			Custom(corerr.ForeignErr).Build()
	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(corterm.Role), role.Name).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, role.Name)
	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)
	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
