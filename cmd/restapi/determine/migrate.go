package determine

import (
	"omega/domain/base/basmodel"
	"omega/internal/core"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine) {
	// Base Domain
	engine.DB.AutoMigrate(&basmodel.BasSetting{})
	engine.DB.AutoMigrate(&basmodel.BasRole{})
	engine.DB.AutoMigrate(&basmodel.BasUser{}).
		AddForeignKey("role_id", "bas_roles(id)", "SET NULL", "SET NULL")
	engine.ActivityDB.AutoMigrate(&basmodel.BasActivity{})

}
