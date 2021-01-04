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
	engine.DB.Exec("ALTER TABLE bas_accounts ADD CONSTRAINT `fk_bas_accounts_self` FOREIGN KEY (parent_id) REFERENCES bas_accounts(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")

	engine.DB.Table(basmodel.UserTable).AutoMigrate(&basmodel.User{})
	// AddForeignKey("role_id", "bas_roles(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("id", "bas_accounts(id)", "RESTRICT", "RESTRICT")
	engine.DB.Exec("ALTER TABLE bas_users ADD CONSTRAINT `fk_bas_users_bas_roles` FOREIGN KEY (role_id) REFERENCES bas_roles(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")
	engine.DB.Exec("ALTER TABLE bas_users ADD CONSTRAINT `fk_bas_users_bas_accounts` FOREIGN KEY (id) REFERENCES bas_accounts(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")

	engine.ActivityDB.Table(basmodel.ActivityTable).AutoMigrate(&basmodel.Activity{})
	engine.DB.Table(basmodel.PhoneTable).AutoMigrate(&basmodel.Phone{})

	engine.DB.Table(basmodel.AccountPhoneTable).AutoMigrate(&basmodel.AccountPhone{})
	// AddForeignKey("account_id", "bas_accounts", "RESTRICT", "RESTRICT").
	// AddForeignKey("phone_id", "bas_phones", "RESTRICT", "RESTRICT")
	engine.DB.Exec("ALTER TABLE bas_account_phones ADD CONSTRAINT `fk_bas_accounts_phones_bas_accounts` FOREIGN KEY (account_id) REFERENCES bas_accounts(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")
	engine.DB.Exec("ALTER TABLE bas_account_phones ADD CONSTRAINT `fk_bas_accounts_phones_bas_phones` FOREIGN KEY (phone_id) REFERENCES bas_phones(id) ON DELETE CASCADE ON UPDATE CASCADE;")

	// EAccounting Domain
	engine.DB.Table(eacmodel.CurrencyTable).AutoMigrate(&eacmodel.Currency{})
	engine.DB.Table(eacmodel.TransactionTable).AutoMigrate(&eacmodel.Transaction{})
	// AddForeignKey("currency_id", "eac_currencies(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("created_by", "bas_users(id)", "RESTRICT", "RESTRICT")
	engine.DB.Exec("ALTER TABLE eac_transactions ADD CONSTRAINT `fk_eac_transactions_eac_currencies` FOREIGN KEY (currency_id) REFERENCES eac_currencies(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")
	engine.DB.Exec("ALTER TABLE eac_transactions ADD CONSTRAINT `fk_eac_transactions_bas_users` FOREIGN KEY (currency_id) REFERENCES bas_users(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")

	engine.DB.Table(eacmodel.SlotTable).AutoMigrate(&eacmodel.Slot{})
	// AddForeignKey("account_id", "bas_accounts(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("transaction_id", "eac_transactions(id)", "RESTRICT", "RESTRICT").
	// AddForeignKey("currency_id", "eac_currencies(id)", "RESTRICT", "RESTRICT")
	engine.DB.Exec("ALTER TABLE eac_slots ADD CONSTRAINT `fk_eac_slot_bas_accounts` FOREIGN KEY (account_id) REFERENCES bas_accounts(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")
	engine.DB.Exec("ALTER TABLE eac_slots ADD CONSTRAINT `fk_eac_slot_eac_transactions` FOREIGN KEY (transaction_id) REFERENCES eac_transactions(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")
	engine.DB.Exec("ALTER TABLE eac_slots ADD CONSTRAINT `fk_eac_slot_eac_currency` FOREIGN KEY (currency_id) REFERENCES eac_currencies(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")

	// Material Domain
	engine.DB.Table(matmodel.CompanyTable).AutoMigrate(&matmodel.Company{})
	engine.DB.Table(matmodel.ColorTable).AutoMigrate(&matmodel.Color{})

	engine.DB.Table(matmodel.GroupTable).AutoMigrate(&matmodel.Group{})
	// AddForeignKey("parent_id", "mat_groups", "RESTRICT", "RESTRICT")
	engine.DB.Exec("ALTER TABLE mat_groups ADD CONSTRAINT `fk_mat_groups_self` FOREIGN KEY (parent_id) REFERENCES mat_groups(id) ON DELETE RESTRICT ON UPDATE RESTRICT;")

	engine.DB.Table(matmodel.UnitTable).AutoMigrate(&matmodel.Unit{})
}
