package coract

// Action is used for type of event and it is shared for all domains
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
