package user

import "omega/internal/model"

// User model
type User struct {
	// gorm.Model
	model.FixedCol
	Name     string      `gorm:"not null;unique" json:"name,omitempty" binding:"required"`
	Username string      `gorm:"not null;unique" json:"username,omitempty" binding:"required"`
	Password string      `gorm:"not null" json:"password,omitempty" binding:"required"`
	Extra    interface{} `sql:"-" json:"extra,omitempty"`
}
