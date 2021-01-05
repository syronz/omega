package eacrepo

import (
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
)

// TransactionRepo for injecting engine
type TransactionRepo struct {
	Engine *core.Engine
	Cols   []string
}

// ProvideTransactionRepo is used in wire and initiate the Cols
func ProvideTransactionRepo(engine *core.Engine) TransactionRepo {
	return TransactionRepo{
		Engine: engine,
		Cols:   helper.TagExtracter(reflect.TypeOf(eacmodel.Transaction{}), eacmodel.TransactionTable),
	}
}

// FindByID finds the transaction via its id
func (p *TransactionRepo) FindByID(fix types.FixedCol) (transaction eacmodel.Transaction, err error) {
	err = p.Engine.ReadDB.Table(eacmodel.TransactionTable).
		Where("company_id = ? AND node_id = ? AND id = ?", fix.CompanyID, fix.NodeID, fix.ID.ToUint64()).
		First(&transaction).Error

	transaction.ID = fix.ID
	err = p.dbError(err, "E1442107", transaction, corterm.List)

	return
}

// List returns an array of transactions
func (p *TransactionRepo) List(params param.Param) (transactions []eacmodel.Transaction, err error) {
	var colsStr string
	if colsStr, err = validator.CheckColumns(p.Cols, params.Select); err != nil {
		err = limberr.Take(err, "E1469040").Build()
		return
	}

	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1494215").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.ReadDB.Table(eacmodel.TransactionTable).Select(colsStr).
		Where(whereStr).
		Order(params.Order).
		Limit(params.Limit).
		Offset(params.Offset).
		Find(&transactions).Error

	err = p.dbError(err, "E1460126", eacmodel.Transaction{}, corterm.List)

	return
}

// Count of transactions, mainly calls with List
func (p *TransactionRepo) Count(params param.Param) (count int64, err error) {
	var whereStr string
	if whereStr, err = params.ParseWhere(p.Cols); err != nil {
		err = limberr.Take(err, "E1436521").Custom(corerr.ValidationFailedErr).Build()
		return
	}

	err = p.Engine.ReadDB.Table(eacmodel.TransactionTable).
		Where(whereStr).
		Count(&count).Error

	err = p.dbError(err, "E1465399", eacmodel.Transaction{}, corterm.List)
	return
}

// Save the transaction, in case it is not exist create it
func (p *TransactionRepo) Save(transaction eacmodel.Transaction) (u eacmodel.Transaction, err error) {
	if err = p.Engine.DB.Table(eacmodel.TransactionTable).Save(&transaction).Error; err != nil {
		err = p.dbError(err, "E1420013", transaction, corterm.Updated)
	}

	p.Engine.DB.Table(eacmodel.TransactionTable).Where("id = ?", transaction.ID).Find(&u)
	return
}

// Create a transaction
func (p *TransactionRepo) Create(transaction eacmodel.Transaction) (u eacmodel.Transaction, err error) {
	if err = p.Engine.DB.Table(eacmodel.TransactionTable).Create(&transaction).Scan(&u).Error; err != nil {
		err = p.dbError(err, "E1413616", transaction, corterm.Created)
	}
	return
}

// Delete the transaction
func (p *TransactionRepo) Delete(transaction eacmodel.Transaction) (err error) {
	if err = p.Engine.DB.Table(eacmodel.TransactionTable).Delete(&transaction).Error; err != nil {
		err = p.dbError(err, "E1474760", transaction, corterm.Deleted)
	}
	return
}

// dbError is an internal method for generate proper dataeace error
func (p *TransactionRepo) dbError(err error, code string, transaction eacmodel.Transaction, action string) error {
	switch corerr.ClearDbErr(err) {
	case corerr.Nil:
		err = nil

	case corerr.NotFoundErr:
		err = corerr.RecordNotFoundHelper(err, code, corterm.ID, transaction.ID, eacterm.Transactions)

	case corerr.ForeignErr:
		err = limberr.Take(err, code).
			Message(corerr.SomeVRelatedToThisVSoItIsNotV, dict.R(basterm.Users),
				dict.R(eacterm.Transaction), dict.R(action)).
			Custom(corerr.ForeignErr).Build()

	case corerr.DuplicateErr:
		err = limberr.Take(err, code).
			Message(corerr.VWithValueVAlreadyExist, dict.R(eacterm.Transaction), transaction.Hash).
			Custom(corerr.DuplicateErr).Build()
		err = limberr.AddInvalidParam(err, "name", corerr.VisAlreadyExist, transaction.Hash)

	case corerr.ValidationFailedErr:
		err = corerr.ValidationFailedHelper(err, code)

	default:
		err = limberr.Take(err, code).
			Message(corerr.InternalServerError).
			Custom(corerr.InternalServerErr).Build()
	}

	return err
}
