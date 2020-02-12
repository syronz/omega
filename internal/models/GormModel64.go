package models

import "time"

// GormModel64 is a replacement for gorm.Model for each table, it has extra fields for sync
type GormModel64 struct {
	ID        uint64     `gorm:"primary_key" json:"id,omitempty" `
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
