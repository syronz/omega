package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/search"
)

// BasActivityRepo for injecting engine
type BasActivityRepo struct {
	Engine *core.Engine
}

// ProvideBasActivityRepo is used in wire
func ProvideBasActivityRepo(engine *core.Engine) BasActivityRepo {
	return BasActivityRepo{Engine: engine}
}

// Save BasActivityRepo
func (p *BasActivityRepo) Save(activity basmodel.BasActivity) (u basmodel.BasActivity, err error) {
	err = p.Engine.ActivityDB.Save(&activity).Error
	return
}

// List of activities
func (p *BasActivityRepo) List(params param.Param) (activities []basmodel.BasActivity, err error) {
	columns, err := basmodel.BasActivity{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.ActivityDB.Select(columns).
		Where(search.Parse(params, basmodel.BasActivity{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&activities).Error

	return
}

// Count of activities
func (p *BasActivityRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.ActivityDB.Table("bas_activities").
		Select(params.Select).
		Where(search.Parse(params, basmodel.BasActivity{}.Pattern())).
		Count(&count).Error
	return
}
