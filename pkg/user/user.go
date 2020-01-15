package user

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name     string      `json:"name"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Extra    interface{} `sql:"-" json:"extra,omitempty"`
}
