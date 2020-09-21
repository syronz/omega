package basmodel

import (
	"omega/internal/core/validator"
	"omega/internal/types"
)

const (
	// ActivityTable is used inside the repo layer
	ActivityTable = "bas_activities"
)

// Activity model
type Activity struct {
	types.FixedCol
	Event    string      `gorm:"index:event_idx" json:"event"`
	UserID   types.RowID `json:"user_id"`
	Username string      `gorm:"index:username_idx" json:"username"`
	IP       string      `json:"ip"`
	URI      string      `json:"uri"`
	Before   string      `gorm:"type:text" json:"before"`
	After    string      `gorm:"type:text" json:"after"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Activity) Pattern() string {
	return `(
		bas_activities.id = '%[1]v' OR
		bas_activities.event LIKE '%[1]v%%' OR
		bas_activities.username LIKE '%[1]v%%' OR
		bas_activities.ip LIKE '%[1]v' OR
		bas_activities.uri LIKE '%[1]v%%' OR
		cast(bas_activities.created_at as char) LIKE '%[1]v%%' OR
		bas_activities.before LIKE '%%%[1]v%%' OR
		bas_activities.after LIKE '%%%[1]v%%'
	)`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Activity) Columns(variate string) (string, error) {
	full := []string{
		"bas_activities.id",
		"bas_activities.event",
		"bas_activities.user_id",
		"bas_activities.username",
		"bas_activities.ip",
		"bas_activities.uri",
		"bas_activities.before",
		"bas_activities.after",
		"bas_activities.created_at",
	}

	return validator.CheckColumns(full, variate)
}
