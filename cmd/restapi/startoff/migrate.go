package startoff

import (
	"omega/domain/base/basmodel"
	"omega/domain/eaccounting/eacmodel"
	"omega/domain/material/matmodel"
	"omega/domain/sync/synmodel"
	"omega/internal/core"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	// Sync Domain
	engine.DB.Table(synmodel.CompanyTable).AutoMigrate(&synmodel.Company{})

	// Base Domain
	engine.DB.Table(basmodel.SettingTable).AutoMigrate(&basmodel.Setting{})
	engine.DB.Exec("ALTER TABLE `bas_settings` ADD UNIQUE `idx_bas_settings_companyID_property`(`company_id`, `property`)")
	engine.DB.Table(basmodel.RoleTable).AutoMigrate(&basmodel.Role{})
	engine.DB.Exec("ALTER TABLE `bas_roles` ADD UNIQUE `idx_bas_roles_companyID_name`(`company_id`, `name`)")

	engine.DB.Table(basmodel.AccountTable).AutoMigrate(&basmodel.Account{})
	// AddForeignKey("parent_id", "bas_accounts(id)", "RESTRICT", "RESTRICT")
	engine.DB.Table(basmodel.UserTable).AutoMigrate(&basmodel.User{})
	// AddForeignKey("role_id", "bas_roles(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("id", "bas_accounts(id)", "RESTRICT", "RESTRICT")
	engine.ActivityDB.Table(basmodel.ActivityTable).AutoMigrate(&basmodel.Activity{})
	engine.DB.Table(basmodel.PhoneTable).AutoMigrate(&basmodel.Phone{})
	engine.DB.Table(basmodel.AccountPhoneTable).AutoMigrate(&basmodel.AccountPhone{})
	// AddForeignKey("account_id", "bas_accounts", "RESTRICT", "RESTRICT").
	// AddForeignKey("phone_id", "bas_phones", "RESTRICT", "RESTRICT")

	// EAccounting Domain
	engine.DB.Table(eacmodel.CurrencyTable).AutoMigrate(&eacmodel.Currency{})
	engine.DB.Table(eacmodel.TransactionTable).AutoMigrate(&eacmodel.Transaction{})
	// AddForeignKey("currency_id", "eac_currencies(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("created_by", "bas_users(id)", "RESTRICT", "RESTRICT")
	engine.DB.Table(eacmodel.SlotTable).AutoMigrate(&eacmodel.Slot{})
	// AddForeignKey("account_id", "bas_accounts(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("transaction_id", "eac_transactions(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("currency_id", "eac_currencies(id)", "RESTRICT", "RESTRICT")

	// Material Domain
	engine.DB.Table(matmodel.CompanyTable).AutoMigrate(&matmodel.Company{})
	engine.DB.Table(matmodel.ColorTable).AutoMigrate(&matmodel.Color{})
	engine.DB.Table(matmodel.GroupTable).AutoMigrate(&matmodel.Group{})
	// AddForeignKey("parent_id", "mat_groups", "RESTRICT", "RESTRICT")
	engine.DB.Table(matmodel.UnitTable).AutoMigrate(&matmodel.Unit{})
}
