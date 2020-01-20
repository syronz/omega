package role

import "omega/internal/models"

// Role model
type Role struct {
	models.FixedCol
	Name        string `gorm:"not null;unique" json:"name,omitempty" binding:"required"`
	Resources   string `gorm:"type:text" json:"resources,omitempty" binding:"-"`
	Description string `json:"description,omitempty" binding:"-"`
}
