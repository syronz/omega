package invoice

import (
	"rest-gin-gorm/pkg/product"
)

type InvoiceDTO struct {
	ID            uint                 `json:"id,string,omitempty"`
	InvoiceNumber string               `json:"invoice_number"`
	Products      []product.ProductDTO `json:"products"`
}
