package table

import (
	"omega/domain/base"
	"omega/domain/base/basmodel"
	"omega/domain/base/basrepo"
	"omega/domain/eaccounting"
	"omega/domain/material"
	"omega/domain/service"
	"omega/domain/sync"
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
			FixedCol: types.FixedCol{
				ID:        1,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name: "Admin",
			Resources: types.ResourceJoin([]types.Resource{
				base.SettingRead, base.SettingWrite, base.SettingExcel,
				base.UserWrite, base.UserRead, base.UserExcel,
				base.ActivitySelf, base.ActivityAll,
				base.RoleRead, base.RoleWrite, base.RoleExcel,
				base.AccountRead, base.AccountWrite, base.AccountExcel,
				eaccounting.CurrencyRead, eaccounting.CurrencyWrite, eaccounting.CurrencyExcel,
				eaccounting.TransactionRead, eaccounting.TransactionManual, eaccounting.TransactionUpdate, eaccounting.TransactionExcel,
				material.CompanyRead, material.CompanyWrite, material.CompanyExcel,
				material.ColorRead, material.ColorWrite, material.ColorExcel,
				material.GroupRead, material.GroupWrite, material.GroupExcel,
			}),
			Description: "admin has all privileges - do not edit",
		},
		{
			FixedCol: types.FixedCol{
				ID:        2,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name: "Cashier",
			Resources: types.ResourceJoin([]types.Resource{
				base.ActivitySelf,
				base.AccountRead, base.AccountWrite, base.AccountExcel,
			}),
			Description: "cashier has privileges for adding transactions - after migration reset",
		},
		{
			FixedCol: types.FixedCol{
				ID:        3,
				CompanyID: engine.Envs.ToUint64(sync.CompanyID),
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name: "Reader",
			Resources: types.ResourceJoin([]types.Resource{
				base.SupperAccess,
				base.SettingRead, base.SettingExcel,
				base.UserRead, base.UserExcel,
				base.RoleRead, base.RoleExcel,
			}),
			Description: "Reader can see all part without changes",
		},
		{
			FixedCol: types.FixedCol{
				ID:        4,
				CompanyID: 1002,
				NodeID:    engine.Envs.ToUint64(sync.NodeID),
			},
			Name: "should_be_deleted",
			Resources: types.ResourceJoin([]types.Resource{
				base.SupperAccess,
				base.SettingRead, base.SettingExcel,
				base.UserRead, base.UserExcel,
				base.RoleRead, base.RoleExcel,
			}),
			Description: "Reader can see all part without changes",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			glog.Fatal("error in saving roles", err)
		}

	}

}
