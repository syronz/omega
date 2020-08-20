package action

// Action is used for type of event
type Action string

// Action enums
const (
	Update Action = "update"
	Create Action = "create"
	Delete Action = "delete"
	Login  Action = "login"
	Save   Action = "save"
	Active Action = "active"
)
