package table

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
)

// InsertRoles for add required roles
func InsertRoles(engine *core.Engine) {
	engine.DB.Exec("UPDATE bas_roles SET deleted_at = null WHERE id IN (1,2,3)")
	roleRepo := basrepo.ProvideRoleRepo(engine)
	roleService := service.ProvideBasRoleService(roleRepo)
	roles := []basmodel.Role{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			Name: "Admin",
			Resources: types.ResourceJoin([]types.Resource{
				base.SettingRead, base.SettingWrite, base.SettingExcel,
				base.UserWrite, base.UserRead, base.UserExcel,
				base.ActivitySelf, base.ActivityAll,
				base.RoleRead, base.RoleWrite, base.RoleExcel,
				// base.Ping,
			}),
			Description: "admin has all privileges - do not edit",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Name: "Cashier",
			Resources: types.ResourceJoin([]types.Resource{
				base.ActivitySelf,
			}),
			Description: "cashier has privileges for adding transactions - after migration reset",
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Name: "Reader",
			Resources: types.ResourceJoin([]types.Resource{
				base.SupperAccess,
				base.SettingRead, base.SettingExcel,
				base.UserRead, base.UserExcel,
				base.RoleRead, base.RoleExcel,
			}),
			Description: "Reade can see all part without changes",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			glog.Fatal("error in saving roles", err)
		}

	}

}
