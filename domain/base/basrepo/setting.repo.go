package basrepo

import (
	"log"
	"omega/domain/base/basmodel"
	"omega/domain/base/message/basterm"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"reflect"

	"github.com/jinzhu/gorm"
)

// SettingRepo for injecting engine
type SettingRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideSettingRepo is used in wire and initiate the Cols
func ProvideSettingRepo(engine *core.Engine) SettingRepo {
	return SettingRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(basmodel.Setting{}), basmodel.SettingTable),
	}
}

// FindByID finds the setting via its id
func (p *SettingRepo) FindByID(fix types.FixedCol) (setting basmodel.Setting, err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&setting).Error

	setting.ID = fix.ID
	err = p.dbError(err, "E1063890", setting, corterm.List)

	return
}

//FindByProperty finds the setting via its property
func (p *SettingRepo) FindByProperty(property string) (setting basmodel.Setting, err error) {
	err = p.Engine.DB.Table(basmodel.SettingTable).
		Select("bas_settings.*").
		Where("property = ?", property).
		Scan(&setting).Error

	log.Printf("found setting : %v", setting)
	//setting.Property = property
	err = p.dbError(err, "E1043108", setting, corterm.List)

	return
}

// List of settings
func (p *SettingRepo) List(params param.Param) (settings []basmodel.Setting, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E1011184").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1039990").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.SettingTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Scan(&settings).Error

	err = p.dbError(err, "E1094986", basmodel.Setting{}, corterm.List)

	return
}

// Count of settings
func (p *SettingRepo) Count(params param.Param) (count uint64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1051896").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.DB.Table(basmodel.SettingTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E1021898", basmodel.Setting{}, corterm.List)
	return
}

// TxCreate is used for creating setting in transaction mode
func (p *SettingRepo) TxCreate(db *gorm.DB, setting basmodel.Setting) (u basmodel.Setting, err error) {
	if err = db.Table(basmodel.SettingTable).Create(&setting).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E1075110", setting, corterm.Created)
	}

	return
}

// Save SettingRepo
func (p *SettingRepo) Save(setting basmodel.Setting) (u basmodel.Setting, err error) {
	if err = p.Engine.DB.Table(basmodel.SettingTable).Save(&setting).Error; err != nil {
		err = p.dbError(err, "E1020662", setting, corterm.Updated)
	}

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

	err = p.dbError(err, "E1081024", basmodel.Setting{}, corterm.Updated)

	p.Engine.DB.Table(basmodel.SettingTable).Where("id = ?", id).Find(&u)
	return
}

// Delete setting
func (p *SettingRepo) Delete(setting basmodel.Setting) (err error) {
	if err = p.Engine.DB.Table(basmodel.SettingTable).Unscoped().Delete(&setting).Error; err != nil {
		err = p.dbError(err, "E1044355", setting, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper database error
func (p *SettingRepo) dbError(err error, code string, setting basmodel.Setting, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, setting.ID, basterm.Settings)

	// case corerr.ForeignErr:
	// 	err = limberr.Take(err, code).
	// 		Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
	// 			dict.R(basterm.Setting), dict.R(action)).
	// 		Custom(corerr.ForeignErr).Build()

	// case corerr.DuplicateErr:
	// 	err = limberr.Take(err, code).
	// 		Message(corerr.VWithValueVAlreadyExist, dict.R(basterm.Setting), setting.Name).
	// 		Custom(corerr.DuplicateErr).Build()
	// 	err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, setting.Name)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
