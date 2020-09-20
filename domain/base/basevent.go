package base

import "omega/internal/types"

// types for base domain
const (
	CreateUser types.Event = "user-create"
	UpdateUser types.Event = "user-update"
	DeleteUser types.Event = "user-delete"
	ListUser   types.Event = "user-list"
	ViewUser   types.Event = "user-view"
	ExcelUser  types.Event = "user-excel"

	CreateRole types.Event = "role-create"
	UpdateRole types.Event = "role-update"
	DeleteRole types.Event = "role-delete"
	ListRole   types.Event = "role-list"
	ViewRole   types.Event = "role-view"
	ExcelRole  types.Event = "role-excel"

	CreateSetting types.Event = "setting-create"
	UpdateSetting types.Event = "setting-update"
	DeleteSetting types.Event = "setting-delete"
	ListSetting   types.Event = "setting-list"
	ViewSetting   types.Event = "setting-view"
	ExcelSetting  types.Event = "setting-excel"

	AllActivity types.Event = "activity-all"

	BasLogin    types.Event = "login"
	BasLogout   types.Event = "logout"
	LoginFailed types.Event = "login-failed"
)
