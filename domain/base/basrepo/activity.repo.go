package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
)

// ActivityRepo for injecting engine
type ActivityRepo struct {
	Engine *core.Engine
}

// ProvideActivityRepo is used in wire
func ProvideActivityRepo(engine *core.Engine) ActivityRepo {
	return ActivityRepo{Engine: engine}
}

// Create ActivityRepo
func (p *ActivityRepo) Create(activity basmodel.Activity) (u basmodel.Activity, err error) {
	err = p.Engine.ActivityDB.Table(basmodel.ActivityTable).Create(&activity).Error
	return
}

// List of activities
func (p *ActivityRepo) List(params param.Param) (activities []basmodel.Activity, err error) {
	columns, err := basmodel.Activity{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.ActivityDB.Select(columns).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&activities).Error

	return
}

// Count of activities
func (p *ActivityRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.ActivityDB.Table("bas_activities").
		Select(params.Select).
		Count(&count).Error
	return
}
