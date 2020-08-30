package basrepo

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/internal/param"
	"omega/internal/search"
	"omega/internal/types"
)

// SettingRepo for injecting engine
type SettingRepo struct {
	Engine *core.Engine
}

// ProvideSettingRepo is used in wire
func ProvideSettingRepo(engine *core.Engine) SettingRepo {
	return SettingRepo{Engine: engine}
}

// FindByID for setting
func (p *SettingRepo) FindByID(id types.RowID) (setting basmodel.Setting, err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).First(&setting, id.ToUint64()).Error
	return
}

// FindByProperty for setting
func (p *SettingRepo) FindByProperty(property string) (setting basmodel.Setting, err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).Where("property = ?", property).First(&setting).Error
	return
}

// List of settings
func (p *SettingRepo) List(params param.Param) (settings []basmodel.Setting, err error) {
	columns, err := basmodel.Setting{}.Columns(params.Select, params)
	if err != nil {
		return
	}

	err = p.Engine.DB.Table(basmodel.SettingTable).Select(columns).
		Where(search.Parse(params, basmodel.Setting{}.Pattern())).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Scan(&settings).Error

	return
}

// Count of settings
func (p *SettingRepo) Count(params param.Param) (count uint64, err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).Table("bas_settings").
		Select(params.Select).
		Where(search.Parse(params, basmodel.Setting{}.Pattern())).
		Count(&count).Error
	return
}

// Save SettingRepo
func (p *SettingRepo) Save(setting basmodel.Setting) (u basmodel.Setting, err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).Save(&setting).Error
	p.Engine.DB.Table(basmodel.SettingTable).Where("id = ?", setting.ID).Find(&u)
	return
}

// Update SettingRepo
func (p *SettingRepo) Update(setting basmodel.Setting) (u basmodel.Setting, err error) {
	id := setting.ID
	setting.ID = 0
	setting.Property = ""
	setting.Type = ""
	setting.Description = ""
	err = p.Engine.DB.Table(basmodel.SettingTable).Where("id = ?", id).Updates(&setting).Error
	p.Engine.DB.Table(basmodel.SettingTable).Where("id = ?", id).Find(&u)
	return
}

// Delete setting
func (p *SettingRepo) Delete(setting basmodel.Setting) (err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).Unscoped().Delete(&setting).Error
	return
}
