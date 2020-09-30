package service

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/corstartoff"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// BasSettingServ for injecting auth basrepo
type BasSettingServ struct {
	Repo   basrepo.SettingRepo
	Engine *core.Engine
}

// ProvideBasSettingService for setting is used in wire
func ProvideBasSettingService(p basrepo.SettingRepo) BasSettingServ {
	return BasSettingServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting setting by it's id
func (p *BasSettingServ) FindByID(id types.RowID) (setting basmodel.Setting, err error) {
	if setting, err = p.Repo.FindByID(id); err != nil {
		err = corerr.Tick(err, "E1064390", "can't fetch the setting", id)
		return
	}

	return
}

// FindByProperty find setting with property
func (p *BasSettingServ) FindByProperty(property string) (setting basmodel.Setting, err error) {
	if setting, err = p.Repo.FindByProperty(property); err != nil {
		err = corerr.Tick(err, "E1026379", "can't fetch the setting by property", property)
		return
	}

	return
}

// List returns setting's property, it support pagination and search and return back count
func (p *BasSettingServ) List(params param.Param) (settings []basmodel.Setting,
	count uint64, err error) {

	if settings, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in users list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in users count")
	}

	return
}

// Save setting
func (p *BasSettingServ) Save(setting basmodel.Setting) (savedSetting basmodel.Setting, err error) {
	if err = setting.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E1066086", "validation failed for saving setting", setting)
		return
	}

	if savedSetting, err = p.Repo.Save(setting); err != nil {
		err = corerr.Tick(err, "E1036118", "error in creating user", setting)
		return
	}

	return
}

// Update setting
func (p *BasSettingServ) Update(setting basmodel.Setting) (savedSetting basmodel.Setting, err error) {
	if err = setting.Validate(coract.Update); err != nil {
		err = corerr.TickValidate(err, "E1053228", "error in updating setting", setting)
		return
	}

	if savedSetting, err = p.Repo.Update(setting); err != nil {
		err = corerr.Tick(err, "E1057541", "setting not updated")
		return
	}

	corstartoff.LoadSetting(p.Engine)

	return
}

// Delete setting, it is soft delete
func (p *BasSettingServ) Delete(settingID types.RowID) (setting basmodel.Setting, err error) {
	if setting, err = p.FindByID(settingID); err != nil {
		err = corerr.Tick(err, "E1076703", "setting not found for deleting, by the way delete setting is forbidden")
		return
	}

	return
}

// Excel is used for export excel file
func (p *BasSettingServ) Excel(params param.Param) (settings []basmodel.Setting, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", basmodel.SettingTable)

	if settings, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E1086162", "cant generate the excel list for setting")
		return
	}

	return
}
