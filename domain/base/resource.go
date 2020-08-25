package base

import "omega/internal/types"

// list of resources for base domain
const (
	SupperAccess types.Resource = "supper:access"

	BasUserNames  types.Resource = "user:names"
	BasUserWrite  types.Resource = "user:write"
	BasUserRead   types.Resource = "user:read"
	BasUserReport types.Resource = "user:report"
	BasUserExcel  types.Resource = "user:excel"

	BasRoleRead  types.Resource = "role:read"
	BasRoleWrite types.Resource = "role:write"
	BasRoleExcel types.Resource = "role:excel"

	BasSettingNames types.Resource = "setting:names"
	BasSettingRead  types.Resource = "setting:read"
	BasSettingWrite types.Resource = "setting:write"
	BasSettingExcel types.Resource = "setting:excel"

	BasActivitySelf types.Resource = "activity:self"
	BasActivityAll  types.Resource = "activity:all"

	BasPing types.Resource = "ping"
)
