package node

import "omega/internal/models"

// Node model
type Node struct {
	models.GormModel64
	CompanyID uint64      `gorm:"index:node_company_idx" json:"company_id"`
	Name      string      `gorm:"not null;unique" json:"name,omitempty" binding:"required"`
	SecretKey string      `json:"secret_key"`
	Phone     string      `gorm:"not null" json:"phone,omitempty"`
	Extra     interface{} `sql:"-" json:"extra,omitempty"`
}
