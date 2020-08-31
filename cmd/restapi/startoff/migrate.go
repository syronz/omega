package startoff

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/internal/core"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	// Base Domain
	engine.DB.Table(basmodel.SettingTable).AutoMigrate(&basmodel.Setting{})
	engine.DB.Table(basmodel.RoleTable).AutoMigrate(&basmodel.Role{})
	engine.DB.Table(basmodel.UserTable).AutoMigrate(&basmodel.User{}).
		AddForeignKey("role_id", fmt.Sprintf("%v(id)", basmodel.RoleTable), "RESTRICT", "RESTRICT")
	engine.ActivityDB.Table(basmodel.ActivityTable).AutoMigrate(&basmodel.Activity{})

}
