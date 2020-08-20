package company

import (
	"fmt"
	"omega/engine"
	"omega/internal/param"
)

// Service for injecting auth repo
type Service struct {
	Repo   Repo
	Engine engine.Engine
}

// ProvideService for company is used in wire
func ProvideService(p Repo) Service {
	return Service{Repo: p, Engine: p.Engine}
}

// FindAll companies
func (p *Service) FindAll() (companies []Company, err error) {
	companies, err = p.Repo.FindAll()
	p.Engine.CheckError(err, "all companies")
	return
}

// List of companies
func (p *Service) List(params param.Param) (data map[string]interface{}, err error) {
	data = make(map[string]interface{})

	data["companies"], err = p.Repo.List(params)
	p.Engine.CheckError(err, "companies list")

	data["count"], err = p.Repo.Count(params)
	p.Engine.CheckError(err, "companies count")

	return
}

// FindByID for company
func (p *Service) FindByID(id uint64) (company Company, err error) {
	company, err = p.Repo.FindByID(id)
	p.Engine.CheckError(err, fmt.Sprintf("Company with id %v", id))

	return
}

// Save company
func (p *Service) Save(company Company) (createdCompany Company, err error) {
	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", company))

	createdCompany, err = p.Repo.Save(company)
	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving company for %+v", company))

	return
}

// func (p *Service) SaveSimple(model interface{}) (createdCompany interface{}, err error) {
// 	company := *(model.(*Company))
// 	company.Password, err = password.Hash(company.Password, p.Engine.Environments.Setting.PasswordSalt)
// 	p.Engine.CheckError(err, fmt.Sprintf("Hashing password failed for %+v", company))

// 	createdCompany, err = p.Repo.Save(company)
// 	p.Engine.CheckInfo(err, fmt.Sprintf("Failed in saving company for %+v", company))

// 	// createdCompany.Password = ""

// 	return
// }

// Delete company
func (p *Service) Delete(company Company) error {
	return p.Repo.Delete(company)
}
