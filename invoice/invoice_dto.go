package invoice

import(
	"rest-gin-gorm/product"
)

type InvoiceDTO struct {
	ID    uint   `json:"id,string,omitempty"`
	InvoiceNumber  string `json:"invoice_number"`
	Products []ProductDTO   `json:"products"`
}
