package basevent

import "omega/internal/types"

// types for base domain
const (
	UserCreate types.Event = "user-create"
	UserUpdate types.Event = "user-update"
	UserDelete types.Event = "user-delete"
	UserList   types.Event = "user-list"
	UserView   types.Event = "user-view"
	UserExcel  types.Event = "user-excel"

	RoleCreate types.Event = "role-create"
	RoleUpdate types.Event = "role-update"
	RoleDelete types.Event = "role-delete"
	RoleList   types.Event = "role-list"
	RoleView   types.Event = "role-view"
	RoleExcel  types.Event = "role-excel"

	SettingCreate types.Event = "setting-create"
	SettingUpdate types.Event = "setting-update"
	SettingDelete types.Event = "setting-delete"
	SettingList   types.Event = "setting-list"
	SettingView   types.Event = "setting-view"
	SettingExcel  types.Event = "setting-excel"

	ActivityAll types.Event = "activity-all"

	BasLogin  types.Event = "login"
	BasLogout types.Event = "logout"
)
