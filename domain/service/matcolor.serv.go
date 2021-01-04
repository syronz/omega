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

// MatColorServ for injecting auth matrepo
type MatColorServ struct {
	Repo   matrepo.ColorRepo
	Engine *core.Engine
}

// ProvideMatColorService for color is used in wire
func ProvideMatColorService(p matrepo.ColorRepo) MatColorServ {
	return MatColorServ{
		Repo:   p,
		Engine: p.Engine,
	}
}

// FindByID for getting color by it's id
func (p *MatColorServ) FindByID(fix types.FixedCol) (color matmodel.Color, err error) {
	if color, err = p.Repo.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E7185082", "can't fetch the color", fix.CompanyID, fix.NodeID, fix.ID)
		return
	}

	return
}

// List of colors, it support pagination and search and return back count
func (p *MatColorServ) List(params param.Param) (colors []matmodel.Color,
	count int64, err error) {

	if params.CompanyID != 0 {
		params.PreCondition = fmt.Sprintf(" mat_colors.company_id = '%v' ", params.CompanyID)
	}

	if colors, err = p.Repo.List(params); err != nil {
		glog.CheckError(err, "error in colors list")
		return
	}

	if count, err = p.Repo.Count(params); err != nil {
		glog.CheckError(err, "error in colors count")
	}

	return
}

// Create a color
func (p *MatColorServ) Create(color matmodel.Color) (createdColor matmodel.Color, err error) {

	if err = color.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E7131519", "validation failed in creating the color", color)
		return
	}

	if createdColor, err = p.Repo.Create(color); err != nil {
		err = corerr.Tick(err, "E7161199", "color not created", color)
		return
	}

	return
}

// Save a color, if it is exist update it, if not create it
func (p *MatColorServ) Save(color matmodel.Color) (savedColor matmodel.Color, err error) {
	if err = color.Validate(coract.Save); err != nil {
		err = corerr.TickValidate(err, "E7143952", corerr.ValidationFailed, color)
		return
	}

	if savedColor, err = p.Repo.Save(color); err != nil {
		err = corerr.Tick(err, "E7173078", "color not saved")
		return
	}

	return
}

// Delete color, it is soft delete
func (p *MatColorServ) Delete(fix types.FixedCol) (color matmodel.Color, err error) {
	if color, err = p.FindByID(fix); err != nil {
		err = corerr.Tick(err, "E7141281", "color not found for deleting")
		return
	}

	if err = p.Repo.Delete(color); err != nil {
		err = corerr.Tick(err, "E7140850", "color not deleted")
		return
	}

	return
}

// Excel is used for export excel file
func (p *MatColorServ) Excel(params param.Param) (colors []matmodel.Color, err error) {
	params.Limit = p.Engine.Envs.ToInt(core.ExcelMaxRows)
	params.Offset = 0
	params.Order = fmt.Sprintf("%v.id ASC", matmodel.ColorTable)

	if colors, err = p.Repo.List(params); err != nil {
		err = corerr.Tick(err, "E7180391", "cant generate the excel list for colors")
		return
	}

	return
}
