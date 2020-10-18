package types

import "time"

type Fixed interface{}

// FixedCol is a replacement for gorm.Model for each table, it has extra fields for sync at
// companyID level
type FixedCol struct {
	ID        RowID      `gorm:"primary_key" json:"id,omitempty"`
	CompanyID uint64     `gorm:"primary_key;auto_increment:false" json:"company_id,omitempty"`
	NodeID    uint64     `json:"node_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// FixedNode is a replacement for gorm.Model for each table, it has extra fields for sync at nodeID
// level
type FixedNode struct {
	ID        RowID      `gorm:"primary_key" json:"id,omitempty"`
	CompanyID uint64     `gorm:"primary_key;auto_increment:false" json:"company_id,omitempty"`
	NodeID    uint64     `gorm:"primary_key;auto_increment:false" json:"node_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
