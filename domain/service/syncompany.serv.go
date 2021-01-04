package service

import (
	"fmt"
	"omega/cmd/restapi/enum/settingfields"
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/sync"
	"omega/domain/sync/synmodel"
	"omega/domain/sync/synrepo"
	"omega/internal/consts"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/glog"

	"github.com/jinzhu/gorm"
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

	db := p.Engine.DB.Begin()

	defer func() {
		if r := recover(); r != nil {
			glog.LogError(fmt.Errorf("panic happened in transaction mode for %v",
				"companies table"), "rollback recover create company")
			db.Rollback()
		}
	}()

	company.Logo = consts.DefaultLogo
	company.Banner = consts.DefaultBanner
	company.Footer = consts.DefaultFooter

	if createdCompany, err = p.Repo.TxCreate(db, company); err != nil {
		err = corerr.Tick(err, "E0910088", "company not created", company)

		db.Rollback()
		return
	}

	// create default roles
	var createdAdminRole basmodel.Role
	if createdAdminRole, err = p.defaultRoles(db, createdCompany.ID.ToUint64(), 101); err != nil {
		err = corerr.Tick(err, "E4219394", "roles not created for company")

		db.Rollback()
		return
	}

	// create admin
	userService := ProvideBasUserService(basrepo.ProvideUserRepo(p.Engine))

	admin := basmodel.User{
		FixedCol: types.FixedCol{
			CompanyID: createdCompany.ID.ToUint64(),
			NodeID:    101,
		},
		RoleID:   createdAdminRole.ID,
		Name:     company.AdminUsername,
		Username: company.AdminUsername,
		Password: company.AdminPassword,
		Lang:     company.Lang,
	}

	if _, err = userService.Create(admin); err != nil {
		err = corerr.Tick(err, "E4294427", "admin not created for company")

		db.Rollback()
		return
	}

	// create default settings
	if err = p.defaultSettings(db, createdCompany, company.Lang); err != nil {
		err = corerr.Tick(err, "E4236774", "settings not created for company")

		db.Rollback()
		return
	}

	db.Commit()
	return
}

func (p *SynCompanyServ) defaultSettings(db *gorm.DB, company synmodel.Company,
	lang dict.Lang) (err error) {
	settingRepo := basrepo.ProvideSettingRepo(p.Engine)
	settingService := ProvideBasSettingService(settingRepo)
	settings := []basmodel.Setting{
		{
			FixedCol: types.FixedCol{
				CompanyID: company.ID.ToUint64(),
				NodeID:    consts.DefaultNodeID,
			},
			Property:    settingfields.CompanyName,
			Value:       company.Name,
			Type:        "string",
			Description: "company's name in the header of invoices",
		},
		{
			FixedCol: types.FixedCol{
				CompanyID: company.ID.ToUint64(),
				NodeID:    consts.DefaultNodeID,
			},
			Property:    settingfields.DefaultLang,
			Value:       string(lang),
			Type:        "string",
			Description: "in case of user JWT not specified this value has been used",
		},
		{
			FixedCol: types.FixedCol{
				CompanyID: company.ID.ToUint64(),
				NodeID:    consts.DefaultNodeID,
			},
			Property:    settingfields.CompanyLogo,
			Value:       consts.DefaultLogo,
			Type:        "string",
			Description: "logo for showed on the application and not invoices",
		},
		{
			FixedCol: types.FixedCol{
				CompanyID: company.ID.ToUint64(),
				NodeID:    consts.DefaultNodeID,
			},
			Property:    settingfields.InvoiceLogo,
			Value:       consts.DefaultLogo,
			Type:        "string",
			Description: "path of logo, if branch logo won’t defined use this logo for invoices",
		},
		{
			FixedCol: types.FixedCol{
				CompanyID: company.ID.ToUint64(),
				NodeID:    consts.DefaultNodeID,
			},
			Property:    settingfields.InvoiceNumberPattern,
			Value:       "location_year_series",
			Type:        "string",
			Description: "location_year_series, location_series, series, year_series, fullyear_series, location_fullyear_series",
		},
	}

	for _, v := range settings {
		if _, err = settingService.TxCreate(db, v); err != nil {
			return
		}
	}

	return
}

// defaultRoles create a roles for specific company
func (p *SynCompanyServ) defaultRoles(db *gorm.DB, companyID,
	nodeID uint64) (createdAdminRole basmodel.Role, err error) {
	roleService := ProvideBasRoleService(basrepo.ProvideRoleRepo(p.Engine))

	adminRole := basmodel.Role{
		FixedCol: types.FixedCol{
			CompanyID: companyID,
			NodeID:    nodeID,
		},
		Name: "Admin",
		Resources: types.ResourceJoin([]types.Resource{
			sync.CompanyRead, sync.CompanyUpdate, sync.CompanyExcel, sync.CompanyRead,
			base.UserWrite, base.UserRead, base.UserExcel, base.RoleRead, base.AccountRead, base.AccountWrite, base.AccountExcel,
			base.SettingRead, base.SettingExcel, base.ActivityCompany, base.ActivitySelf, base.PhoneRead, base.PhoneWrite,
		}),
		Description: "Admin has all privileges per the company",
	}

	if createdAdminRole, err = roleService.Create(adminRole); err != nil {
		return
	}

	roles := []basmodel.Role{
		{
			FixedCol: types.FixedCol{
				CompanyID: companyID,
				NodeID:    nodeID,
			},
			Name: "Reader",
			Resources: types.ResourceJoin([]types.Resource{
				sync.CompanyRead, sync.CompanyExcel,
				base.UserRead, base.UserExcel, base.RoleRead, base.AccountRead,
				base.AccountExcel, base.ActivityCompany, base.ActivitySelf, base.PhoneRead,
			}),
			Description: "Reader can see all part without changes",
		},
		{
			FixedCol: types.FixedCol{
				CompanyID: companyID,
				NodeID:    nodeID,
			},
			Name: "Cashier",
			Resources: types.ResourceJoin([]types.Resource{
				sync.CompanyRead, sync.CompanyExcel,
				base.UserRead, base.UserExcel, base.AccountWrite, base.AccountExcel, base.ActivitySelf, base.PhoneRead, base.PhoneWrite,
			}),
			Description: "Cashier can create, append the invoice, also have access to delete uninserted results",
		},
	}

	for _, v := range roles {
		if _, err = roleService.TxCreate(db, v); err != nil {
			return
		}
	}

	return
}

// Save a company, if it is exist update it, if not create it
func (p *SynCompanyServ) Save(company synmodel.Company) (savedCompany synmodel.Company, err error) {
	// TODO we have change coract.create to coract.save
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

// UploadImage is used to save the path of the new picture of company table
func (p *SynCompanyServ) UploadImage(company synmodel.Company, imageType string) (updatedImage synmodel.Company, err error) {
	if err = company.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E0945980", corerr.ValidationFailed, company)
		return
	}
	if updatedImage, err = p.Repo.UpdateImage(company, imageType); err != nil {
		err = corerr.Tick(err, "E0903417", "update logo company failed")
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
