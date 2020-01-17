package model

import "time"

// FixedCol is a replacement for gorm.Model for each table, it has extra fields for sync
type FixedCol struct {
	ID        uint64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
