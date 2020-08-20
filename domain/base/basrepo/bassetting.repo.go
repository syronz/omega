package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/search"
	"omega/internal/types"
)

// BasSettingRepo for injecting engine
type BasSettingRepo struct {
	Engine *core.Engine
}

// ProvideBasSettingRepo is used in wire
func ProvideBasSettingRepo(engine *core.Engine) BasSettingRepo {
	return BasSettingRepo{Engine: engine}
}

// FindByID for setting
func (p *BasSettingRepo) FindByID(id types.RowID) (setting basmodel.BasSetting, err error) {
	err = p.Engine.DB.First(&setting, id.ToUint64()).Error
	return
}

// FindByProperty for setting
func (p *BasSettingRepo) FindByProperty(property string) (setting basmodel.BasSetting, err error) {
	err = p.Engine.DB.Where("property = ?", property).First(&setting).Error
	return
}

// List of settings
func (p *BasSettingRepo) List(params param.Param) (settings []basmodel.BasSetting, err error) {
	columns, err := basmodel.BasSetting{}.Columns(params.Select)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table("bas_settings").Select(columns).
		Where(search.Parse(params, basmodel.BasSetting{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Scan(&settings).Error

	return
}

// Count of settings
func (p *BasSettingRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table("bas_settings").
		Select(params.Select).
		Where(search.Parse(params, basmodel.BasSetting{}.Pattern())).
		Count(&count).Error
	return
}

// Save BasSettingRepo
func (p *BasSettingRepo) Save(setting basmodel.BasSetting) (u basmodel.BasSetting, err error) {
	err = p.Engine.DB.Save(&setting).Error
	p.Engine.DB.Where("id = ?", setting.ID).Find(&u)
	return
}

// Update BasSettingRepo
func (p *BasSettingRepo) Update(setting basmodel.BasSetting) (u basmodel.BasSetting, err error) {
	setting.Property = ""
	setting.Type = ""
	setting.Description = ""
	err = p.Engine.DB.Model(&setting).Updates(&setting).Error
	p.Engine.DB.Where("id = ?", setting.ID).Find(&u)
	return
}

// Delete setting
func (p *BasSettingRepo) Delete(setting basmodel.BasSetting) (err error) {
	err = p.Engine.DB.Unscoped().Delete(&setting).Error
	return
}
