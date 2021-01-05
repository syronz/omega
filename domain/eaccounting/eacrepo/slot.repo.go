package eacrepo

import (
	"errors"
	"omega/domain/base/message/basterm"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/eaccounting/eacterm"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"reflect"

	"gorm.io/gorm"
)

// SlotRepo for injecting engine
type SlotRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideSlotRepo is used in wire and initiate the Cols
func ProvideSlotRepo(engine *core.Engine) SlotRepo {
	return SlotRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(eacmodel.Slot{}), eacmodel.SlotTable),
	}
}

// FindByID finds the slot via its id
func (p *SlotRepo) FindByID(fix types.FixedCol) (slot eacmodel.Slot, err error) {
	err = p.Engine.ReadDB.Table(eacmodel.SlotTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&slot).Error

	slot.ID = fix.ID
	err = p.dbError(err, "E1471037", slot, corterm.List)

	return
}

// List returns an array of slots
func (p *SlotRepo) List(params param.Param) (slots []eacmodel.Slot, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E1471933").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1490185").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.ReadDB.Table(eacmodel.SlotTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&slots).Error

	err = p.dbError(err, "E1430518", eacmodel.Slot{}, corterm.List)

	return
}

// Count of slots, mainly calls with List
func (p *SlotRepo) Count(params param.Param) (count int64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1428251").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.ReadDB.Table(eacmodel.SlotTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E1481282", eacmodel.Slot{}, corterm.List)
	return
}

// Save the slot, in case it is not exist create it
func (p *SlotRepo) Save(slot eacmodel.Slot) (u eacmodel.Slot, err error) {
	if err = p.Engine.DB.Table(eacmodel.SlotTable).Save(&slot).Error; err != nil {
		err = p.dbError(err, "E1484537", slot, corterm.Updated)
	}

	p.Engine.DB.Table(eacmodel.SlotTable).Where("id = ?", slot.ID).Find(&u)
	return
}

// Create a slot
func (p *SlotRepo) Create(slot eacmodel.Slot) (u eacmodel.Slot, err error) {
	if err = p.Engine.DB.Table(eacmodel.SlotTable).Create(&slot).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E1437304", slot, corterm.Created)
	}
	return
}

// Delete the slot
func (p *SlotRepo) Delete(slot eacmodel.Slot) (err error) {
	if err = p.Engine.DB.Table(eacmodel.SlotTable).Delete(&slot).Error; err != nil {
		err = p.dbError(err, "E1494853", slot, corterm.Deleted)
	}
	return
}

// LastSlot returns the last slot before post_date
func (p *SlotRepo) LastSlot(slotIn eacmodel.Slot) (slot eacmodel.Slot, err error) {
	err = p.Engine.ReadDB.Table(eacmodel.SlotTable).
		Where("company_id = ? AND account_id = ? AND currency_id = ? AND post_date <= ?",
			slotIn.CompanyID, slotIn.AccountID, slotIn.CurrencyID, slotIn.PostDate).
		Order(" post_date DESC, id DESC ").
		Limit(1).
		Find(&slot).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	err = p.dbError(err, "E1471037", slot, corterm.List)

	return
}

// LastSlotWithID returns the last slot before post_date
func (p *SlotRepo) LastSlotWithID(slotIn eacmodel.Slot) (slot eacmodel.Slot, err error) {
	err = p.Engine.ReadDB.Table(eacmodel.SlotTable).
		Where("company_id = ? AND account_id = ? AND currency_id = ? AND post_date <= ? AND id < ?",
			slotIn.CompanyID, slotIn.AccountID, slotIn.CurrencyID, slotIn.PostDate, slotIn.ID).
		Order(" post_date DESC, id DESC ").
		Limit(1).
		Find(&slot).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}

	err = p.dbError(err, "E1495037", slot, corterm.List)

	return
}

// RegulateBalances will adjust all balances after the post_date
func (p *SlotRepo) RegulateBalances(slotIn eacmodel.Slot, adjust float64) (err error) {
	err = p.Engine.DB.Table(eacmodel.SlotTable).
		Where("company_id = ? AND account_id = ? AND currency_id = ? AND post_date > ?",
			slotIn.CompanyID, slotIn.AccountID, slotIn.CurrencyID, slotIn.PostDate).
		Update("balance", gorm.Expr("balance + ?", adjust)).Error
	return
}

// RegulateBalancesSave will adjust all balances after the post_date
func (p *SlotRepo) RegulateBalancesSave(slotIn eacmodel.Slot, adjust float64) (err error) {
	err = p.Engine.DB.Table(eacmodel.SlotTable).
		Where("company_id = ? AND account_id = ? AND currency_id = ? AND post_date >= ? AND id > ?",
			slotIn.CompanyID, slotIn.AccountID, slotIn.CurrencyID, slotIn.PostDate, slotIn.ID).
		Update("balance", gorm.Expr("balance + ?", adjust)).Error
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *SlotRepo) dbError(err error, code string, slot eacmodel.Slot, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, slot.ID, eacterm.Transaction)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(eacterm.Transaction), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
