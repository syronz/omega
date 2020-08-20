package basevent

import "omega/internal/types"

// types for base domain
const (
	BasUserCreate types.Event = "user-create"
	BasUserUpdate types.Event = "user-update"
	BasUserDelete types.Event = "user-delete"
	BasUserList   types.Event = "user-list"
	BasUserView   types.Event = "user-view"
	BasUserExcel  types.Event = "user-excel"

	BasRoleCreate types.Event = "role-create"
	BasRoleUpdate types.Event = "role-update"
	BasRoleDelete types.Event = "role-delete"
	BasRoleList   types.Event = "role-list"
	BasRoleView   types.Event = "role-view"
	BasRoleExcel  types.Event = "role-excel"

	BasSettingCreate types.Event = "setting-create"
	BasSettingUpdate types.Event = "setting-update"
	BasSettingDelete types.Event = "setting-delete"
	BasSettingList   types.Event = "setting-list"
	BasSettingView   types.Event = "setting-view"
	BasSettingExcel  types.Event = "setting-excel"

	BasActivityAll types.Event = "activity-all"

	BasLogin  types.Event = "login"
	BasLogout types.Event = "logout"
)
