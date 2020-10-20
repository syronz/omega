package service

import (
	"fmt"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// EacCurrencyServ for injecting auth eacrepo
type EacCurrencyServ struct {
	Repo   eacrepo.CurrencyRepo
	Engine *core.Engine
}

// ProvideEacCurrencyService for currency is used in wire
func ProvideEacCurrencyService(p eacrepo.CurrencyRepo) EacCurrencyServ {
	return EacCurrencyServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting currency by it's id
func (p *EacCurrencyServ) FindByID(fix types.FixedCol) (currency eacmodel.Currency, err error) {
	if currency, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1432098", "can't fetch the currency", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	return
}

// List of currencies, it support pagination and search and return back count
func (p *EacCurrencyServ) List(params param.Param) (currencies []eacmodel.Currency,
	count uint64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" eac_currencies.company_id = '%v' ", params.CompanyID)
	}

	if currencies, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in currencies list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in currencies count")
	}

	return
}

// Create a currency
func (p *EacCurrencyServ) Create(currency eacmodel.Currency) (createdCurrency eacmodel.Currency, err error) {

	if err = currency.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1493520", "validation failed in creating the currency", currency)
		return
	}

	if createdCurrency, err = p.Repo.Create(currency); err != nil {
		err = corerr.Tick(err, "E1478011", "currency not created", currency)
		return
	}

	return
}

// Save a currency, if it is exist update it, if not create it
func (p *EacCurrencyServ) Save(currency eacmodel.Currency) (savedCurrency eacmodel.Currency, err error) {
	if err = currency.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1431170", corerr.ValidationFailed, currency)
		return
	}

	if savedCurrency, err = p.Repo.Save(currency); err != nil {
		err = corerr.Tick(err, "E1496019", "currency not saved")
		return
	}

	return
}

// Delete currency, it is soft delete
func (p *EacCurrencyServ) Delete(fix types.FixedCol) (currency eacmodel.Currency, err error) {
	if currency, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E1457804", "currency not found for deleting")
		return
	}

	if err = p.Repo.Delete(currency); err != nil {
		err = corerr.Tick(err, "E1428337", "currency not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *EacCurrencyServ) Excel(params param.Param) (currencies []eacmodel.Currency, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", eacmodel.CurrencyTable)

	if currencies, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1443095", "cant generate the excel list for currencies")
		return
	}

	return
}
