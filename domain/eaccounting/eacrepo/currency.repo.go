package eacrepo

import (
	"omega/domain/base/message/basterm"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacterm"
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

// CurrencyRepo for injecting engine
type CurrencyRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideCurrencyRepo is used in wire and initiate the Cols
func ProvideCurrencyRepo(engine *core.Engine) CurrencyRepo {
	return CurrencyRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(eacmodel.Currency{}), eacmodel.CurrencyTable),
	}
}

// FindByID finds the currency via its id
func (p *CurrencyRepo) FindByID(fix types.FixedCol) (currency eacmodel.Currency, err error) {
	err = p.Engine.DB.Table(eacmodel.CurrencyTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&currency).Error

	currency.ID = fix.ID
	err = p.dbError(err, "E1420401", currency, corterm.List)

	return
}

// List returns an array of currencies
func (p *CurrencyRepo) List(params param.Param) (currencies []eacmodel.Currency, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E1412501").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1492582").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(eacmodel.CurrencyTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&currencies).Error

	err = p.dbError(err, "E1422926", eacmodel.Currency{}, corterm.List)

	return
}

// Count of currencies, mainly calls with List
func (p *CurrencyRepo) Count(params param.Param) (count int64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1439729").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(eacmodel.CurrencyTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E1491513", eacmodel.Currency{}, corterm.List)
	return
}

// Save the currency, in case it is not exist create it
func (p *CurrencyRepo) Save(currency eacmodel.Currency) (u eacmodel.Currency, err error) {
	if err = p.Engine.DB.Table(eacmodel.CurrencyTable).Save(&currency).Error; err != nil {
		err = p.dbError(err, "E1426565", currency, corterm.Updated)
	}

	p.Engine.DB.Table(eacmodel.CurrencyTable).Where("id = ?", currency.ID).Find(&u)
	return
}

// Create a currency
func (p *CurrencyRepo) Create(currency eacmodel.Currency) (u eacmodel.Currency, err error) {
	if err = p.Engine.DB.Table(eacmodel.CurrencyTable).Create(&currency).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E1475182", currency, corterm.Created)
	}
	return
}

// Delete the currency
func (p *CurrencyRepo) Delete(currency eacmodel.Currency) (err error) {
	if err = p.Engine.DB.Table(eacmodel.CurrencyTable).Delete(&currency).Error; err != nil {
		err = p.dbError(err, "E1426280", currency, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *CurrencyRepo) dbError(err error, code string, currency eacmodel.Currency, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, currency.ID, eacterm.Currencies)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(eacterm.Currency), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(eacterm.Currency), currency.Name).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, currency.Name)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
