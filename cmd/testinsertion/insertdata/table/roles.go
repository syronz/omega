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
	roleRepo := basrepo.ProvideRoleRepo(engine)
	roleService := service.ProvideBasRoleService(roleRepo)

	// reset the roles table
	roleRepo.Engine.DB.Table(basmodel.RoleTable).Unscoped().Delete(basmodel.Role{})

	roles := []basmodel.Role{
		{
			GormCol: types.GormCol{
				ID: 1,
			},
			Name: "Super-Admin",
			Resources: types.ResourceJoin([]types.Resource{
				base.SettingRead, base.SettingWrite, base.SettingExcel,
				base.UserWrite, base.UserRead, base.UserExcel,
				base.ActivitySelf, base.ActivityAll,
				base.RoleRead, base.RoleWrite, base.RoleExcel,
			}),
			Description: "super-admin has all privileges - do not edit",
		},
		{
			GormCol: types.GormCol{
				ID: 2,
			},
			Name: "Admin",
			Resources: types.ResourceJoin([]types.Resource{
				base.SettingRead, base.SettingWrite, base.SettingExcel,
				base.UserWrite, base.UserRead, base.UserExcel,
				base.ActivitySelf, base.ActivityAll,
				base.RoleRead, base.RoleWrite, base.RoleExcel,
			}),
			Description: "admin has all privileges - do not edit",
		},
		{
			GormCol: types.GormCol{
				ID: 3,
			},
			Name:        "Cashier",
			Resources:   types.ResourceJoin([]types.Resource{base.ActivitySelf}),
			Description: "cashier has all privileges - after migration reset",
		},
		{
			GormCol: types.GormCol{
				ID: 4,
			},
			Name:        "for foreign 1",
			Resources:   string(base.SettingRead),
			Description: "for foreign 1",
		},
		{
			GormCol: types.GormCol{
				ID: 5,
			},
			Name:        "for update 1",
			Resources:   string(base.SettingRead),
			Description: "for update 1",
		},
		{
			GormCol: types.GormCol{
				ID: 6,
			},
			Name:        "for update 2",
			Resources:   string(base.SettingRead),
			Description: "for update 2",
		},
		{
			GormCol: types.GormCol{
				ID: 7,
			},
			Name:        "for delete 1",
			Resources:   string(base.SettingRead),
			Description: "for delete 1",
		},
		{
			GormCol: types.GormCol{
				ID: 8,
			},
			Name:        "for search 1",
			Resources:   string(base.SettingRead),
			Description: "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 9,
			},
			Name:        "for search 2",
			Resources:   string(base.SettingRead),
			Description: "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 10,
			},
			Name:        "for search 3",
			Resources:   string(base.SettingRead),
			Description: "searchTerm1",
		},
		{
			GormCol: types.GormCol{
				ID: 11,
			},
			Name:        "for delete 2",
			Resources:   string(base.SettingRead),
			Description: "for delete 2",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			glog.Fatal(err)
		}

	}

}
