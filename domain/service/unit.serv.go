package service

import (
	"fmt"
	"omega/domain/material/matmodel"
	"omega/domain/material/matrepo"
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/glog"
)

// MatUnitServ for injecting auth matrepo
type MatUnitServ struct {
	Repo   matrepo.UnitRepo
	Engine *core.Engine
}

// ProvideMatUnitService for unit is used in wire
func ProvideMatUnitService(p matrepo.UnitRepo) MatUnitServ {
	return MatUnitServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting unit by it's id
func (p *MatUnitServ) FindByID(fix types.FixedCol) (unit matmodel.Unit, err error) {
	if unit, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E7126006", "can't fetch the unit", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	return
}

// List of units, it support pagination and search and return back count
func (p *MatUnitServ) List(params param.Param) (units []matmodel.Unit,
	count uint64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" mat_units.company_id = '%v' ", params.CompanyID)
	}

	if units, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in units list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in units count")
	}

	return
}

// Create a unit
func (p *MatUnitServ) Create(unit matmodel.Unit) (createdUnit matmodel.Unit, err error) {

	if err = unit.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E7151496", "validation failed in creating the unit", unit)
		return
	}

	if createdUnit, err = p.Repo.Create(unit); err != nil {
		err = corerr.Tick(err, "E7126648", "unit not created", unit)
		return
	}

	return
}

// Save a unit, if it is exist update it, if not create it
func (p *MatUnitServ) Save(unit matmodel.Unit) (savedUnit matmodel.Unit, err error) {
	if err = unit.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E7186722", corerr.ValidationFailed, unit)
		return
	}

	if savedUnit, err = p.Repo.Save(unit); err != nil {
		err = corerr.Tick(err, "E7169808", "unit not saved")
		return
	}

	return
}

// Delete unit, it is soft delete
func (p *MatUnitServ) Delete(fix types.FixedCol) (unit matmodel.Unit, err error) {
	if unit, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E7114397", "unit not found for deleting")
		return
	}

	if err = p.Repo.Delete(unit); err != nil {
		err = corerr.Tick(err, "E7147933", "unit not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *MatUnitServ) Excel(params param.Param) (units []matmodel.Unit, err error) {
	params.Limit = p.Engine.Envs.ToUint64(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", matmodel.UnitTable)

	if units, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E7110296", "cant generate the excel list for units")
		return
	}

	return
}
