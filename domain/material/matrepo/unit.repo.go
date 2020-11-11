package matrepo

import (
	"omega/domain/base/message/basterm"
	"omega/domain/material/matmodel"
	"omega/domain/material/matterm"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"reflect"
)

// UnitRepo for injecting engine
type UnitRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideUnitRepo is used in wire and initiate the Cols
func ProvideUnitRepo(engine *core.Engine) UnitRepo {
	return UnitRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(matmodel.Unit{}), matmodel.UnitTable),
	}
}

// FindByID finds the unit via its id
func (p *UnitRepo) FindByID(fix types.FixedCol) (unit matmodel.Unit, err error) {
	err = p.Engine.DB.Table(matmodel.UnitTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&unit).Error

	unit.ID = fix.ID
	err = p.dbError(err, "E7168159", unit, corterm.List)

	return
}

// List returns an array of units
func (p *UnitRepo) List(params param.Param) (units []matmodel.Unit, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E7149588").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E7150516").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(matmodel.UnitTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&units).Error

	err = p.dbError(err, "E7135200", matmodel.Unit{}, corterm.List)

	return
}

// Count of units, mainly calls with List
func (p *UnitRepo) Count(params param.Param) (count uint64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E7160160").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(matmodel.UnitTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E7140882", matmodel.Unit{}, corterm.List)
	return
}

// Save the unit, in case it is not exist create it
func (p *UnitRepo) Save(unit matmodel.Unit) (u matmodel.Unit, err error) {
	if err = p.Engine.DB.Table(matmodel.UnitTable).Save(&unit).Error; err != nil {
		err = p.dbError(err, "E7126064", unit, corterm.Updated)
	}

	p.Engine.DB.Table(matmodel.UnitTable).Where("id = ?", unit.ID).Find(&u)
	return
}

// Create a unit
func (p *UnitRepo) Create(unit matmodel.Unit) (u matmodel.Unit, err error) {
	if err = p.Engine.DB.Table(matmodel.UnitTable).Create(&unit).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E7129940", unit, corterm.Created)
	}
	return
}

// Delete the unit
func (p *UnitRepo) Delete(unit matmodel.Unit) (err error) {
	if err = p.Engine.DB.Table(matmodel.UnitTable).Delete(&unit).Error; err != nil {
		err = p.dbError(err, "E7198048", unit, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *UnitRepo) dbError(err error, code string, unit matmodel.Unit, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, unit.ID, matterm.Units)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(matterm.Unit), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(matterm.Unit), unit.Name).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, unit.Name)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
