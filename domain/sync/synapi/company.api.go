package synapi

import (
	"net/http"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/domain/sync"
	"omega/domain/sync/synmodel"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"

	"github.com/gin-gonic/gin"
)

// CompanyAPI for injecting company service
type CompanyAPI struct {
	Service service.SynCompanyServ
	Engine  *core.Engine
}

// ProvideCompanyAPI for company is used in wire
func ProvideCompanyAPI(c service.SynCompanyServ) CompanyAPI {
	return CompanyAPI{Service: c, Engine: c.Engine}
}

// FindByID is used for fetch a company by it's id
func (p *CompanyAPI) FindByID(c *gin.Context) {
	resp := response.New(p.Engine, c, sync.Domain)
	var err error
	var company synmodel.Company
	id, err := types.StrToRowID(c.Param("companyID"))
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	if company, err = p.Service.FindByID(id); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(sync.ViewCompany)
	resp.Status(http.StatusOK).
		MessageT(corterm.VInfo, basterm.Company).
		JSON(company)
}

// List of companies
func (p *CompanyAPI) List(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, synmodel.CompanyTable, sync.Domain)

	data := make(map[string]interface{})
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E0937019"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	if data["list"], data["count"], err = p.Service.List(params); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(sync.ListCompany)
	resp.Status(http.StatusOK).
		MessageT(corterm.ListOfV, basterm.Companies).
		JSON(data)
}

// Create company
func (p *CompanyAPI) Create(c *gin.Context) {
	resp := response.New(p.Engine, c, sync.Domain)
	var company, createdCompany synmodel.Company
	var err error

	if err = resp.Bind(&company, "E0929072", sync.Domain, basterm.Company); err != nil {
		return
	}

	if createdCompany, err = p.Service.Create(company); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.RecordCreate(sync.CreateCompany, company)
	resp.Status(http.StatusOK).
		MessageT(corterm.VCreatedSuccessfully, basterm.Company).
		JSON(createdCompany)
}

// Update company
func (p *CompanyAPI) Update(c *gin.Context) {
	resp := response.New(p.Engine, c, sync.Domain)
	var err error

	var company, companyBefore, companyUpdated synmodel.Company

	id, err := types.StrToRowID(c.Param("companyID"))
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	if companyBefore, err = p.Service.FindByID(id); err != nil {
		resp.Error(err).JSON()
		return
	}

	if companyUpdated, err = p.Service.Save(company); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(sync.UpdateCompany, companyBefore, company)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Company).
		JSON(companyUpdated)
}

// Delete company
func (p *CompanyAPI) Delete(c *gin.Context) {
	resp := response.New(p.Engine, c, sync.Domain)
	var err error
	var company synmodel.Company

	id, err := types.StrToRowID(c.Param("companyID"))
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	if company, err = p.Service.Delete(id); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(sync.DeleteCompany, company)
	resp.Status(http.StatusOK).
		MessageT(corterm.VDeletedSuccessfully, basterm.Company).
		JSON()
}

// Excel generate excel files eaced on search
func (p *CompanyAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Companies, sync.Domain)
	var err error

	if params.CompanyID, err = resp.GetCompanyID("E0966407"); err != nil {
		return
	}

	if !resp.CheckRange(params.CompanyID, 0) {
		return
	}

	companies, err := p.Service.Excel(params)
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	ex := excel.New("company")
	ex.AddSheet("Companies").
		AddSheet("Summary").
		Active("Companies").
		SetPageLayout("landscape", "A4").
		SetPageMargins(0.2).
		SetHeaderFooter().
		SetColWidth("B", "F", 15.3).
		SetColWidth("G", "G", 40).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Companies").
		WriteHeader("ID", "Company ID", "Node ID", "Name", "Symbol", "Code", "Updated At").
		SetSheetFields("ID", "CompanyID", "NodeID", "Name", "Symbol", "Code", "UpdatedAt").
		WriteData(companies).
		AddTable()

	buffer, downloadName, err := ex.Generate()
	if err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(sync.ExcelCompany)

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", buffer.Bytes())

}
