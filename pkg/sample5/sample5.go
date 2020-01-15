package sample5

import "github.com/jinzhu/gorm"

type Sample5 struct {
	gorm.Model
	Code  string
	Price uint
	Name  string
	Extra interface{} `sql:"-" json:"Extra,omitempty"`
}
