package base

import "omega/internal/types"

// list of resources for base domain
const (
	Domain string = "base"

	SupperAccess types.Resource = "supper:access"

	UserWrite types.Resource = "user:write"
	UserRead  types.Resource = "user:read"
	UserExcel types.Resource = "user:excel"

	RoleRead  types.Resource = "role:read"
	RoleWrite types.Resource = "role:write"
	RoleExcel types.Resource = "role:excel"

	AccountRead  types.Resource = "account:read"
	AccountWrite types.Resource = "account:write"
	AccountExcel types.Resource = "account:excel"

	SettingRead  types.Resource = "setting:read"
	SettingWrite types.Resource = "setting:write"
	SettingExcel types.Resource = "setting:excel"

	//for activity all companies we use SuperAdmin resource
	ActivityCompany types.Resource = "activity:company"
	ActivitySelf    types.Resource = "activity:self"

	PhoneRead  types.Resource = "phone:read"
	PhoneWrite types.Resource = "phone:write"
	PhoneExcel types.Resource = "phone:excel"

	Ping types.Resource = "ping"
)
