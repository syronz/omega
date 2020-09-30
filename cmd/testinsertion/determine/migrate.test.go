package determine

import (
	"fmt"
	"omega/domain/base/basmodel"
	"omega/internal/core"
	"omega/pkg/glog"
)

// Migrate the database for creating tables
func Migrate(engine *core.Engine, noReset bool) {

	if !noReset {
		dropTable(engine)
	}

	// Base Domain
	err := engine.DB.Table(basmodel.SettingTable).AutoMigrate(&basmodel.Setting{}).Error
	glog.CheckError(err, "error in migrating settings")
	err = engine.DB.Table(basmodel.RoleTable).AutoMigrate(&basmodel.Role{}).Error
	glog.CheckError(err, "error in migrating roles")
	err = engine.DB.Table(basmodel.UserTable).AutoMigrate(&basmodel.User{}).
		AddForeignKey("role_id", fmt.Sprintf("%v(id)", basmodel.RoleTable), "RESTRICT", "RESTRICT").Error
	glog.CheckError(err, "error in migrating users")
	err = engine.ActivityDB.Table(basmodel.ActivityTable).AutoMigrate(&basmodel.Activity{}).Error
	glog.CheckError(err, "error in migrating activities")

}

func dropTable(engine *core.Engine) {
	var err error
	if err = engine.DB.DropTable(basmodel.UserTable).Error; err != nil {
		glog.Error(fmt.Sprintf("error in dropping %v table", basmodel.UserTable), err)
	}
	if err = engine.DB.DropTable(basmodel.RoleTable).Error; err != nil {
		glog.Error(fmt.Sprintf("error in dropping %v table", basmodel.RoleTable), err)
	}
}
