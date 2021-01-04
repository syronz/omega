package service

import (
	"fmt"
	"omega/domain/material/matmodel"
	"omega/domain/material/matrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// MatCompanyServ for injecting auth matrepo
type MatCompanyServ struct {
	Repo   matrepo.CompanyRepo
	Engine *core.Engine
}

// ProvideMatCompanyService for company is used in wire
func ProvideMatCompanyService(p matrepo.CompanyRepo) MatCompanyServ {
	return MatCompanyServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting company by it's id
func (p *MatCompanyServ) FindByID(fix types.FixedCol) (company matmodel.Company, err error) {
	if company, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E7121746", "can't fetch the company", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	return
}

// List of companies, it support pagination and search and return back count
func (p *MatCompanyServ) List(params param.Param) (companies []matmodel.Company,
	count int64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" mat_companies.company_id = '%v' ", params.CompanyID)
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
func (p *MatCompanyServ) Create(company matmodel.Company) (createdCompany matmodel.Company, err error) {

	if err = company.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E7129096", "validation failed in creating the company", company)
		return
	}

	if createdCompany, err = p.Repo.Create(company); err != nil {
		err = corerr.Tick(err, "E7110088", "company not created", company)
		return
	}

	return
}

// Save a company, if it is exist update it, if not create it
func (p *MatCompanyServ) Save(company matmodel.Company) (savedCompany matmodel.Company, err error) {
	if err = company.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E7137980", corerr.ValidationFailed, company)
		return
	}

	if savedCompany, err = p.Repo.Save(company); err != nil {
		err = corerr.Tick(err, "E7145417", "company not saved")
		return
	}

	return
}

// Delete company, it is soft delete
func (p *MatCompanyServ) Delete(fix types.FixedCol) (company matmodel.Company, err error) {
	if company, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E7199162", "company not found for deleting")
		return
	}

	if err = p.Repo.Delete(company); err != nil {
		err = corerr.Tick(err, "E7194293", "company not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *MatCompanyServ) Excel(params param.Param) (companies []matmodel.Company, err error) {
	params.Limit = p.Engine.Envs.ToInt(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", matmodel.CompanyTable)

	if companies, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E7150013", "cant generate the excel list for companies")
		return
	}

	return
}
