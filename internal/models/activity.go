package models

// import "omega/internal/models"

// Activity model
type Activity struct {
	FixedCol
	Event    string `gorm:"index:event_idx" json:"event"`
	UserID   uint64 `json:"user_id"`
	Username string `gorm:"index:username_idx" json:"username"`
	IP       string `json:"ip"`
	URI      string `json:"uri"`
	Before   string `gorm:"type:text" json:"before"`
	After    string `gorm:"type:text" json:"after"`
}
