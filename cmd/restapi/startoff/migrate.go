package startoff

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/material/matmodel"
	"omega/internal/core"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	// Base Domain
	engine.DB.Table(basmodel.SettingTable).AutoMigrate(&basmodel.Setting{})
	engine.DB.Table(basmodel.RoleTable).AutoMigrate(&basmodel.Role{})
	engine.DB.Table(basmodel.AccountTable).AutoMigrate(&basmodel.Account{}).
		AddForeignKey("parent_id", "bas_accounts(id)", "RESTRICT", "RESTRICT")
	engine.DB.Table(basmodel.UserTable).AutoMigrate(&basmodel.User{}).
		AddForeignKey("role_id", fmt.Sprintf("%v(id)", basmodel.RoleTable), "RESTRICT", "RESTRICT").
		AddForeignKey("id", "bas_accounts(id)", "RESTRICT", "RESTRICT")
	engine.ActivityDB.Table(basmodel.ActivityTable).AutoMigrate(&basmodel.Activity{})

	// EAccounting Domain
	engine.DB.Table(eacmodel.CurrencyTable).AutoMigrate(&eacmodel.Currency{})
	engine.DB.Table(eacmodel.TransactionTable).AutoMigrate(&eacmodel.Transaction{}).
		AddForeignKey("currency_id", "eac_currencies(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("created_by", "bas_users(id)", "RESTRICT", "RESTRICT")
	engine.DB.Table(eacmodel.SlotTable).AutoMigrate(&eacmodel.Slot{}).
		AddForeignKey("account_id", "bas_accounts(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("transaction_id", "eac_transactions(id)", "RESTRICT", "RESTRICT").
		AddForeignKey("currency_id", "eac_currencies(id)", "RESTRICT", "RESTRICT")

	// Material Domain
	engine.DB.Table(matmodel.CompanyTable).AutoMigrate(&matmodel.Company{})
	engine.DB.Table(matmodel.ColorTable).AutoMigrate(&matmodel.Color{})
}
