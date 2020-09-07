package service

import (
	"fmt"
	"net/http"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
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
	return BasSettingServ{Repo: p, Engine: p.Engine}
}

// FindByID for getting setting by it's id
func (p *BasSettingServ) FindByID(id types.RowID) (setting basmodel.Setting, err error) {
	setting, err = p.Repo.FindByID(id)
	glog.CheckInfo(err, fmt.Sprintf("Setting with id %v", id))

	return
}

// FindByProperty find setting with property
func (p *BasSettingServ) FindByProperty(property string) (setting basmodel.Setting, err error) {
	setting, err = p.Repo.FindByProperty(property)
	glog.CheckError(err, fmt.Sprintf("Setting with property %v", property))

	return
}

// List returns setting's property, it support pagination and search and return back count
func (p *BasSettingServ) List(params param.Param) (data map[string]interface{}, err error) {

	data = make(map[string]interface{})

	params.Pagination.Limit = 100
	params.Pagination.Order = "id asc"

	data["list"], err = p.Repo.List(params)
	glog.CheckError(err, "settings list")
	if err != nil {
		return
	}

	data["count"], err = p.Repo.Count(params)
	glog.CheckError(err, "settings count")

	return
}

// Save setting
func (p *BasSettingServ) Save(setting basmodel.Setting, params param.Param) (savedSetting basmodel.Setting, err error) {
	if err = setting.Validate(coract.Save); err != nil {
		glog.CheckError(err, "Error in saving setting")
		return
	}

	savedSetting, err = p.Repo.Save(setting)

	return
}

// Update setting
func (p *BasSettingServ) Update(setting basmodel.Setting, params param.Param) (savedSetting basmodel.Setting, err error) {
	if err = setting.Validate(coract.Update); err != nil {
		in this part I should add custom error with new part
		glog.CheckError(err, "Error in saving setting")
		return
	}

	savedSetting, err = p.Repo.Update(setting)
	corstartoff.LoadSetting(p.Engine)

	return
}

// Delete setting, it is soft delete
func (p *BasSettingServ) Delete(settingID types.RowID) (setting basmodel.Setting, err error) {
	if setting, err = p.FindByID(settingID); err != nil {
		return setting, core.NewErrorWithStatus(err.Error(), http.StatusNotFound)
	}

	return
}

// Excel is used for export excel file
func (p *BasSettingServ) Excel(params param.Param) (settings []basmodel.Setting, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = "bas_settings.id ASC"

	settings, err = p.Repo.List(params)
	glog.CheckError(err, "settings excel")

	return
}
