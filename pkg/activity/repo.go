package activity

import (
	"omega/engine"
	"omega/internal/models"
	"omega/internal/param"
	"omega/utils/search"
)

// Repo for injecting engine
type Repo struct {
	Engine engine.Engine
}

var pattern = `(activities.event LIKE '%%%[1]v%%' OR
		activities.username LIKE '%[1]v' OR
		activities.id LIKE '%[1]v' OR
		activities.created_at LIKE '%[1]v%%' OR
		activities.data LIKE '%%%[1]v%%' OR
		activities.ip LIKE '%[1]v')`

// ProvideRepo is used in wire
func ProvideRepo(engine engine.Engine) Repo {
	return Repo{Engine: engine}
}

// List activities
func (p *Repo) List(params param.Param) (activities []models.Activity, err error) {
	err = p.Engine.DB.Select(params.Select).
		Where(search.Parse(params, pattern)).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&activities).Error

	return
}

// Count activities
func (p *Repo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("activities").
		Select(params.Select).
		Where(search.Parse(params, pattern)).
		Count(&count).Error
	return
}

// FindByID for activity
func (p *Repo) FindByID(id uint64) (activity models.Activity, err error) {
	err = p.Engine.DB.First(&activity, id).Error

	return
}
