package invoice

import "github.com/jinzhu/gorm"

type Invoice struct {
	gorm.Model
	InvoiceNumber  string
}