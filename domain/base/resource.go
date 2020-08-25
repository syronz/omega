package base

import "omega/internal/types"

// list of resources for base domain
const (
	SupperAccess types.Resource = "supper:access"

	UserNames  types.Resource = "user:names"
	UserWrite  types.Resource = "user:write"
	UserRead   types.Resource = "user:read"
	UserReport types.Resource = "user:report"
	UserExcel  types.Resource = "user:excel"

	RoleRead  types.Resource = "role:read"
	RoleWrite types.Resource = "role:write"
	RoleExcel types.Resource = "role:excel"

	SettingNames types.Resource = "setting:names"
	SettingRead  types.Resource = "setting:read"
	SettingWrite types.Resource = "setting:write"
	SettingExcel types.Resource = "setting:excel"

	ActivitySelf types.Resource = "activity:self"
	ActivityAll  types.Resource = "activity:all"

	Ping types.Resource = "ping"
)
