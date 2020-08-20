package company

import "omega/internal/models"

// Company model
type Company struct {
	models.GormModel64
	Name       string      `gorm:"not null;unique" json:"name,omitempty"`
	key        string      `json:"key"`
	Expiration string      `json:"expiration,omitempty"`
	Plan       string      `gorm:"not null" json:"plan,omitempty"`
	Detail     string      `gorm:"not null" json:"detail,omitempty"`
	Phone      string      `gorm:"not null" json:"phone,omitempty"`
	Extra      interface{} `sql:"-" json:"extra,omitempty"`
}
