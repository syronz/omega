package sample4

import "github.com/jinzhu/gorm"

type Sample4 struct {
	gorm.Model
	Code  string
	Price uint
	Name  string
	Extra interface{} `sql:"-" json:"Extra,omitempty"`
}