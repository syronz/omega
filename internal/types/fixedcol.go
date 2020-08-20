package types

import "time"

// FixedCol is a replacement for gorm.Model for each table, it has extra fields for sync
type FixedCol struct {
	ID        RowID      `gorm:"primary_key" json:"id,omitempty" `
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
