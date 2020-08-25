package table

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/types"
)

// InsertBasRoles for add required roles
func InsertBasRoles(engine *core.Engine) {
	engine.DB.Exec("UPDATE bas_roles SET deleted_at = null WHERE id IN (1,2,3)")
	roleRepo := basrepo.ProvideBasRoleRepo(engine)
	roleService := service.ProvideBasRoleService(roleRepo)
	roles := []basmodel.BasRole{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			Name: "Admin",
			Resources: types.ResourceJoin([]types.Resource{
				// base.BasSettingRead, base.BasSettingWrite, base.BasSettingExcel,
				base.BasUserNames, base.BasUserWrite, base.BasUserRead, base.BasUserReport, base.BasUserExcel,
				base.BasActivitySelf, base.BasActivityAll,
				base.BasRoleRead, base.BasRoleWrite, base.BasRoleExcel,
				// base.BasPing,
			}),
			Description: "admin has all privileges - do not edit",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Name: "Cashier",
			Resources: types.ResourceJoin([]types.Resource{
				base.BasActivitySelf,
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
				base.BasSettingRead, base.BasSettingExcel,
				base.BasUserNames, base.BasUserRead, base.BasUserReport, base.BasUserExcel,
				base.BasRoleRead, base.BasRoleExcel,
			}),
			Description: "Reade can see all part without changes",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
