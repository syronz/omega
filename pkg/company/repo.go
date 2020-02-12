package company

import (
	"omega/engine"
	"omega/internal/param"
	"omega/utils/search"
)

// Repo for injecting engine
type Repo struct {
	Engine engine.Engine
}

var pattern = `(companies.name LIKE '%[1]v%%' OR
		companies.id = '%[1]v' OR
		companies.plan LIKE '%[1]v' OR
		companies.phone LIKE '%[1]v%%')`

// ProvideRepo is used in wire
func ProvideRepo(engine engine.Engine) Repo {
	return Repo{Engine: engine}
}

// FindAll companies
func (p *Repo) FindAll() (companies []Company, err error) {
	err = p.Engine.DB.Select("id, name").Find(&companies).Error
	return
}

// List companies
func (p *Repo) List(params param.Param) (companies []Company, err error) {
	err = p.Engine.DB.Select(params.Select).
		Where(search.Parse(params, pattern)).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&companies).Error

	return
}

// Count companies
func (p *Repo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("companies").
		Select(params.Select).
		Where("deleted_at = null").
		Where(search.Parse(params, pattern)).
		Count(&count).Error
	return
}

// FindByID for company
func (p *Repo) FindByID(id uint64) (company Company, err error) {
	err = p.Engine.DB.First(&company, id).Error

	return
}

// FindByCompanyname for company
func (p *Repo) FindByCompanyname(companyname string) (company Company, err error) {
	err = p.Engine.DB.Where("companyname = ?", companyname).First(&company).Error
	return
}

// Save company
func (p *Repo) Save(company Company) (u Company, err error) {
	err = p.Engine.DB.Save(&company).Scan(&u).Error
	return
}

// Delete company
func (p *Repo) Delete(company Company) (err error) {
	err = p.Engine.DB.Delete(&company).Error
	return
}
