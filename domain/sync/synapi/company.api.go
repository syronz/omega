package synapi

import (
	"fmt"
	"net/http"
	"omega/domain/base/message/basterm"
	"omega/domain/service"
	"omega/domain/sync"
	"omega/domain/sync/synmodel"
	"omega/domain/sync/synterm"
	"omega/internal/core"
	"omega/internal/core/corterm"
	"omega/internal/response"
	"omega/internal/types"
	"omega/pkg/excel"
	"omega/pkg/random"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

	if !resp.CheckRange(params.CompanyID) {
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

	// glog.Debug(company)

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

	if err = resp.Bind(&company, "E098732", sync.Domain, synterm.Company); err != nil {
		return
	}

	if companyBefore, err = p.Service.FindByID(id); err != nil {
		resp.Error(err).JSON()
		return
	}

	company.ID = id
	company.Logo = companyBefore.Logo
	company.Banner = companyBefore.Banner
	company.Footer = companyBefore.Footer
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

//Upload Allow end user to upload the image of the company
func (p *CompanyAPI) Upload(c *gin.Context) {
	resp := response.New(p.Engine, c, sync.Domain)
	var err error

	var company, companyBefore, updatedImage synmodel.Company

	id, _ := types.StrToRowID(c.Param("companyID"))

	if companyBefore, err = p.Service.FindByID(id); err != nil {
		resp.Error(err).JSON()
		return
	}

	company = companyBefore

	dirName := c.Param("dirName")
	var folderName string
	var filename string

	if dirName == "logo" {
		folderName = p.Engine.Envs[sync.CompanyLogo]
		filename = company.Logo
	} else if dirName == "banner" {
		folderName = p.Engine.Envs[sync.CompanyBanner]
		filename = company.Banner
	} else if dirName == "footer" {
		folderName = p.Engine.Envs[sync.CompanyFooter]
		filename = company.Footer
	} else {
		resp.Error(err).JSON()
		return
	}

	file, err := c.FormFile("picture")
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}
	maxImageSize, _ := strconv.ParseInt(p.Engine.Envs[sync.MaxImageSize], 10, 64)
	if file.Size > maxImageSize {
		c.String(http.StatusBadRequest, "Image is too large")
		return
	} else if !(file.Header["Content-Type"][0] == "image/jpeg" ||
		file.Header["Content-Type"][0] == "image/png") {
		c.String(http.StatusBadRequest, "Uploaded file is not an image")
		return
	}

	os.MkdirAll(folderName, os.ModePerm)
	oldFile := filepath.Join(folderName, filename)
	if strings.Contains(oldFile, "-") {
		os.Remove(oldFile)
	}

	fileExt := filepath.Ext(file.Filename)
	newFileName := random.String(20)
	newFileName = fmt.Sprintf(`%v-%v%v`, id, newFileName, fileExt)
	newFilePath := fmt.Sprintf(`%v/%v`, folderName, newFileName)
	if err := c.SaveUploadedFile(file, newFilePath); err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
		return
	}

	if dirName == "logo" {
		company.Logo = newFileName
	} else if dirName == "banner" {
		company.Banner = newFileName
	} else if dirName == "footer" {
		company.Footer = newFileName
	}
	if updatedImage, err = p.Service.Save(company); err != nil {
		resp.Error(err).JSON()
		return
	}

	resp.Record(sync.UpdateCompany, companyBefore, company)
	resp.Status(http.StatusOK).
		MessageT(corterm.VUpdatedSuccessfully, basterm.Company).
		JSON(updatedImage)
}

//Download Allow end user to view the image of the company
func (p *CompanyAPI) Download(c *gin.Context) {
	// resp := response.New(p.Engine, c, sync.Domain)
	filename := c.Param("filename")
	dirName := c.Param("dirName")
	// id, err := strconv.ParseUint(c.Param("companyID"), 10, 16)
	// if err != nil {
	// 	resp.Error(err).JSON()
	// 	return
	// }

	var folderName string
	if dirName == "logo" {
		folderName = p.Engine.Envs[sync.CompanyLogo]
	} else if dirName == "banner" {
		folderName = p.Engine.Envs[sync.CompanyBanner]
	} else if dirName == "footer" {
		folderName = p.Engine.Envs[sync.CompanyFooter]
	}

	var fileFullPath string
	fileFullPath = filepath.Join(folderName, filename)
	c.FileAttachment(fileFullPath, filename)
}

// Excel generate excel files eaced on search
func (p *CompanyAPI) Excel(c *gin.Context) {
	resp, params := response.NewParam(p.Engine, c, basterm.Companies, sync.Domain)
	var err error

	// if params.CompanyID, err = resp.GetCompanyID("E0966407"); err != nil {
	// 	return
	// }

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
		SetColWidth("G", "G", 25).
		SetColWidth("H", "N", 15.3).
		SetColWidth("O", "O", 35).
		Active("Summary").
		SetColWidth("A", "D", 20).
		Active("Companies").
		WriteHeader("ID", "Name", "Legal Name", "Key", "ServerAddress", "Expiration",
			"License", "Plan", "Detail", "Phone", "Email", "Website", "Type", "Code", "Updated At").
		SetSheetFields("ID", "Name", "LegalName", "Key", "ServerAddress", "Expiration",
			"License", "Plan", "Detail", "Phone", "Email", "Website", "Type", "Code", "UpdatedAt").
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
