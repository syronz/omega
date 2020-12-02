package service

import (
	"fmt"
	"omega/domain/sync/synmodel"
	"omega/domain/sync/synrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// SynCompanyServ for injecting auth synrepo
type SynCompanyServ struct {
	Repo   synrepo.CompanyRepo
	Engine *core.Engine
}

// ProvideSynCompanyService for company is used in wire
func ProvideSynCompanyService(p synrepo.CompanyRepo) SynCompanyServ {
	return SynCompanyServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting company by it's id
func (p *SynCompanyServ) FindByID(id types.RowID) (company synmodel.Company, err error) {
	if company, err = p.Repo.FindByID(id); err != nil {
		err = corerr.Tick(err, "E0921746", "can't fetch the company", id)
		return
	}

	return
}

// List of companies, it support pagination and search and return back count
func (p *SynCompanyServ) List(params param.Param) (companies []synmodel.Company,
	count uint64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" syn_companies.company_id = '%v' ", params.CompanyID)
	}

	if companies, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in companies list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in companies count")
	}

	return
}

// Create a company
func (p *SynCompanyServ) Create(company synmodel.Company) (createdCompany synmodel.Company, err error) {

	if err = company.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E0929096", "validation failed in creating the company", company)
		return
	}

	if createdCompany, err = p.Repo.Create(company); err != nil {
		err = corerr.Tick(err, "E0910088", "company not created", company)
		return
	}

	return
}

// Save a company, if it is exist update it, if not create it
func (p *SynCompanyServ) Save(company synmodel.Company) (savedCompany synmodel.Company, err error) {
	if err = company.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E0937980", corerr.ValidationFailed, company)
		return
	}

	if savedCompany, err = p.Repo.Save(company); err != nil {
		err = corerr.Tick(err, "E0945417", "company not saved")
		return
	}

	return
}

// Delete company, it is soft delete
func (p *SynCompanyServ) Delete(id types.RowID) (company synmodel.Company, err error) {
	if company, err = p.FindByID(id); err != nil {
		err = corerr.Tick(err, "E0999162", "company not found for deleting")
		return
	}

	if err = p.Repo.Delete(company); err != nil {
		err = corerr.Tick(err, "E0994293", "company not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *SynCompanyServ) Excel(params param.Param) (companies []synmodel.Company, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", synmodel.CompanyTable)

	if companies, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E0950013", "cant generate the excel list for companies")
		return
	}

	return
}
