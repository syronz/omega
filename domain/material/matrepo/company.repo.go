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

// CompanyRepo for injecting engine
type CompanyRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideCompanyRepo is used in wire and initiate the Cols
func ProvideCompanyRepo(engine *core.Engine) CompanyRepo {
	return CompanyRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(matmodel.Company{}), matmodel.CompanyTable),
	}
}

// FindByID finds the company via its id
func (p *CompanyRepo) FindByID(fix types.FixedCol) (company matmodel.Company, err error) {
	err = p.Engine.DB.Table(matmodel.CompanyTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&company).Error

	company.ID = fix.ID
	err = p.dbError(err, "E7190154", company, corterm.List)

	return
}

// List returns an array of companies
func (p *CompanyRepo) List(params param.Param) (companies []matmodel.Company, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E7124397").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E7148082").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(matmodel.CompanyTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&companies).Error

	err = p.dbError(err, "E7126210", matmodel.Company{}, corterm.List)

	return
}

// Count of companies, mainly calls with List
func (p *CompanyRepo) Count(params param.Param) (count int64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E7159547").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(matmodel.CompanyTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E7118234", matmodel.Company{}, corterm.List)
	return
}

// Save the company, in case it is not exist create it
func (p *CompanyRepo) Save(company matmodel.Company) (u matmodel.Company, err error) {
	if err = p.Engine.DB.Table(matmodel.CompanyTable).Save(&company).Error; err != nil {
		err = p.dbError(err, "E7169136", company, corterm.Updated)
	}

	p.Engine.DB.Table(matmodel.CompanyTable).Where("id = ?", company.ID).Find(&u)
	return
}

// Create a company
func (p *CompanyRepo) Create(company matmodel.Company) (u matmodel.Company, err error) {
	if err = p.Engine.DB.Table(matmodel.CompanyTable).Create(&company).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E7118953", company, corterm.Created)
	}
	return
}

// Delete the company
func (p *CompanyRepo) Delete(company matmodel.Company) (err error) {
	if err = p.Engine.DB.Table(matmodel.CompanyTable).Delete(&company).Error; err != nil {
		err = p.dbError(err, "E7154485", company, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *CompanyRepo) dbError(err error, code string, company matmodel.Company, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, company.ID, basterm.Companies)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(basterm.Company), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(basterm.Company), company.Name).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, company.Name)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
