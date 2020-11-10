package matrepo

import (
	"omega/domain/base/message/basterm"
	"omega/domain/material/matmodel"
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

// ColorRepo for injecting engine
type ColorRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideColorRepo is used in wire and initiate the Cols
func ProvideColorRepo(engine *core.Engine) ColorRepo {
	return ColorRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(matmodel.Color{}), matmodel.ColorTable),
	}
}

// FindByID finds the color via its id
func (p *ColorRepo) FindByID(fix types.FixedCol) (color matmodel.Color, err error) {
	err = p.Engine.DB.Table(matmodel.ColorTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&color).Error

	color.ID = fix.ID
	err = p.dbError(err, "E7175881", color, corterm.List)

	return
}

// List returns an array of colors
func (p *ColorRepo) List(params param.Param) (colors []matmodel.Color, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E7185994").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E7115460").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(matmodel.ColorTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&colors).Error

	err = p.dbError(err, "E7171752", matmodel.Color{}, corterm.List)

	return
}

// Count of colors, mainly calls with List
func (p *ColorRepo) Count(params param.Param) (count uint64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E7131878").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(matmodel.ColorTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E7152018", matmodel.Color{}, corterm.List)
	return
}

// Save the color, in case it is not exist create it
func (p *ColorRepo) Save(color matmodel.Color) (u matmodel.Color, err error) {
	if err = p.Engine.DB.Table(matmodel.ColorTable).Save(&color).Error; err != nil {
		err = p.dbError(err, "E7181350", color, corterm.Updated)
	}

	p.Engine.DB.Table(matmodel.ColorTable).Where("id = ?", color.ID).Find(&u)
	return
}

// Create a color
func (p *ColorRepo) Create(color matmodel.Color) (u matmodel.Color, err error) {
	if err = p.Engine.DB.Table(matmodel.ColorTable).Create(&color).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E7191565", color, corterm.Created)
	}
	return
}

// Delete the color
func (p *ColorRepo) Delete(color matmodel.Color) (err error) {
	if err = p.Engine.DB.Table(matmodel.ColorTable).Delete(&color).Error; err != nil {
		err = p.dbError(err, "E7123790", color, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *ColorRepo) dbError(err error, code string, color matmodel.Color, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, color.ID, basterm.Colors)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(basterm.Color), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(basterm.Color), color.Name).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, color.Name)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
