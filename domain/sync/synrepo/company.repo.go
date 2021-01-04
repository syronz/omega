package synrepo

import (
	"omega/domain/base/message/basterm"
	"omega/domain/sync/synmodel"
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

	"github.com/jinzhu/gorm"
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
		Cols:   helper.TagExtracter(reflect.TypeOf(synmodel.Company{}), synmodel.CompanyTable),
	}
}

// FindByID finds the company via its id
func (p *CompanyRepo) FindByID(id types.RowID) (company synmodel.Company, err error) {
	err = p.Engine.DB.Table(synmodel.CompanyTable).
		Where("id = ? ", id).
		First(&company).Error

	company.ID = id
	err = p.dbError(err, "E0990154", company, corterm.List)

	return
}

// List returns an array of companies
func (p *CompanyRepo) List(params param.Param) (companies []synmodel.Company, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E0924397").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E0948082").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(synmodel.CompanyTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&companies).Error

	err = p.dbError(err, "E0926210", synmodel.Company{}, corterm.List)

	return
}

// Count of companies, mainly calls with List
func (p *CompanyRepo) Count(params param.Param) (count uint64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E0959547").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(synmodel.CompanyTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E0918234", synmodel.Company{}, corterm.List)
	return
}

// Save the company, in case it is not exist create it
func (p *CompanyRepo) Save(company synmodel.Company) (u synmodel.Company, err error) {
	if err = p.Engine.DB.Table(synmodel.CompanyTable).Save(&company).Error; err != nil {
		err = p.dbError(err, "E0969136", company, corterm.Updated)
	}

	p.Engine.DB.Table(synmodel.CompanyTable).Where("id = ?", company.ID).Find(&u)
	return
}

// TxCreate a company
func (p *CompanyRepo) TxCreate(db *gorm.DB, company synmodel.Company) (u synmodel.Company, err error) {
	if err = db.Table(synmodel.CompanyTable).Create(&company).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E0918953", company, corterm.Created)
	}
	return
}

// Delete the company
func (p *CompanyRepo) Delete(company synmodel.Company) (err error) {
	if err = p.Engine.DB.Table(synmodel.CompanyTable).Delete(&company).Error; err != nil {
		err = p.dbError(err, "E0954485", company, corterm.Deleted)
	}
	return
}

// save the path of the new picture of company table
func (p *CompanyRepo) UpdateImage(company synmodel.Company, imageType string) (u synmodel.Company, err error) {
	var imageUpdate synmodel.Company

	if imageType == "logo" {
		imageUpdate = synmodel.Company{Logo: company.Logo}
	} else if imageType == "banner" {
		imageUpdate = synmodel.Company{Banner: company.Banner}
	} else if imageType == "footer" {
		imageUpdate = synmodel.Company{Footer: company.Footer}
	}
	err = p.Engine.DB.Model(&company).Updates(imageUpdate).Scan(&u).Error
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *CompanyRepo) dbError(err error, code string, company synmodel.Company, action string) error {
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
		err = limberr.AddInvalidParam(err, "name, license", corerr.VisAlreadyExist, company.Name)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
