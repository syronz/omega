package migrate

import (
	"omega/engine"
	"omega/internal/models"
	"omega/pkg/role"
)

func migrateRoles(e engine.Engine) {
	roleRepo := role.ProvideRepo(e)
	roleService := role.ProvideService(roleRepo)
	roles := []role.Role{
		{
			FixedCol: models.FixedCol{
				ID: 1,
			},
			Name:        "Admin",
			Resources:   "users:read users:write users:report activities:self activities:all roles:read roles:write",
			Description: "admin has all privileges - do not edit",
		},
		{
			FixedCol: models.FixedCol{
				ID: 2,
			},
			Name:        "Cashier",
			Resources:   "invoices:read invoices:write",
			Description: "cashier has all privileges - after migration reset",
		},
	}

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			e.ServerLog.Fatal(err)
		}

	}

}
