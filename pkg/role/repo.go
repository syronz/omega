package role

import (
	"omega/engine"
	"omega/internal/param"
	"omega/utils/search"
)

// Repo for injecting engine
type Repo struct {
	Engine engine.Engine
}

var pattern = `(roles.name LIKE '%%%[1]v%%' OR
		roles.resources LIKE '%%%[1]v%%')`

// ProvideRepo is used in wire
func ProvideRepo(engine engine.Engine) Repo {
	return Repo{Engine: engine}
}

// FindAll roles
func (p *Repo) FindAll() (roles []Role, err error) {
	err = p.Engine.DB.Select("id, name").Find(&roles).Error
	return
}

// List roles
func (p *Repo) List(params param.Param) (roles []Role, err error) {
	err = p.Engine.DB.Select(params.Select).
		Where(search.Parse(params, pattern)).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&roles).Error

	return
}

// Count roles
func (p *Repo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("roles").
		Select(params.Select).
		Where("deleted_at = null").
		Where(search.Parse(params, pattern)).
		Count(&count).Error
	return
}

// FindByID for role
func (p *Repo) FindByID(id uint64) (role Role, err error) {
	err = p.Engine.DB.First(&role, id).Error

	return
}

// Save role
func (p *Repo) Save(role Role) (u Role, err error) {
	err = p.Engine.DB.Save(&role).Scan(&u).Error
	return
}

// Delete role
func (p *Repo) Delete(role Role) (err error) {
	err = p.Engine.DB.Delete(&role).Error
	return
}
