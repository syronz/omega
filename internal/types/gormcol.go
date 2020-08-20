package types

import "time"

// GormCol is a same as model.gorm, we use our name if in futer customize it don't face problem
type GormCol struct {
	ID        RowID      `gorm:"primary_key" json:"id,omitempty" `
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
