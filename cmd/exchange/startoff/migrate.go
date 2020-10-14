package startoff

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/domain/eaccounting/eacmodel"
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

}
