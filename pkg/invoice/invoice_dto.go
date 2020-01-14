package invoice

import (
	"omega/pkg/product"
)

type InvoiceDTO struct {
	ID            uint                 `json:"id,string,omitempty"`
	InvoiceNumber string               `json:"invoice_number"`
	Products      []product.ProductDTO `json:"products"`
}
