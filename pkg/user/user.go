package user

import (
	"omega/internal/models"
)

// User model
type User struct {
	models.FixedCol
	RoleID   uint64      `gorm:"index:role_id_idx" json:"role_id"`
	Name     string      `gorm:"not null;unique" json:"name,omitempty" binding:"required"`
	Username string      `gorm:"not null;unique" json:"username,omitempty" binding:"required"`
	Password string      `gorm:"not null" json:"password,omitempty" binding:"-"`
	Phone    string      `gorm:"not null" json:"phone,omitempty"`
	Extra    interface{} `sql:"-" json:"extra,omitempty"`
}
