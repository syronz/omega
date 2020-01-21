package migrate

import (
	"omega/engine"
	// "omega/internal/models"
	"omega/pkg/role"
)

func migrateRoles(e engine.Engine) {
	roleRepo := role.ProvideRepo(e)
	roleService := role.ProvideService(roleRepo)
	roles := []role.Role{
		{
			Name:        "Admin",
			Resources:   "user:read user:write user:report activities:self activities:all roles:read roles:write",
			Description: "admin has all privileges - do not edit",
		},
		{
			Name:        "Cashier",
			Resources:   "invoices:read invoices:write",
			Description: "cashier has all privileges - after migration reset",
		},
	}
	roles[0].ID = 1
	roles[1].ID = 2

	for _, v := range roles {
		if _, err := roleService.Save(v); err != nil {
			e.ServerLog.Fatal(err)
		}

	}

}
