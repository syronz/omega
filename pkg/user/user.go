package user

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Name     string      `json:"name,omitempty"`
	Username string      `json:"username,omitempty"`
	Password string      `json:"password,omitempty"`
	Extra    interface{} `sql:"-" json:"extra,omitempty"`
}
