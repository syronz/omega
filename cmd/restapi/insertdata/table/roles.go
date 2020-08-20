package table

import (
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/base/basresource"
	"omega/domain/service"
	"omega/internal/core"
	"omega/internal/types"
	"strings"
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
			Resources: strings.Join([]string{
				basresource.BasSettingRead, basresource.BasSettingWrite, basresource.BasSettingExcel,
				basresource.BasUserNames, basresource.BasUserWrite, basresource.BasUserRead, basresource.BasUserReport, basresource.BasUserExcel,
				basresource.BasActivitySelf, basresource.BasActivityAll,
				basresource.BasRoleRead, basresource.BasRoleWrite, basresource.BasRoleExcel,
			}, ", "),
			Description: "admin has all privileges - do not edit",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Name: "Cashier",
			Resources: strings.Join([]string{
				basresource.BasActivitySelf,
			}, ", "),
			Description: "cashier has privileges for adding transactions - after migration reset",
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Name: "Reader",
			Resources: strings.Join([]string{
				basresource.SupperAccess,
				basresource.BasSettingRead, basresource.BasSettingExcel,
				basresource.BasUserNames, basresource.BasUserRead, basresource.BasUserReport, basresource.BasUserExcel,
				basresource.BasRoleRead, basresource.BasRoleExcel,
			}, ", "),
			Description: "Reade can see all part without changes",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			engine.ServerLog.Fatal(err)
		}

	}

}
