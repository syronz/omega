package invoice

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

func ProvideInvoiceRepostiory(DB *gorm.DB) InvoiceRepository {
	return InvoiceRepository{DB: DB}
}

func (p *InvoiceRepository) Save(invoice Invoice) Invoice {
	p.DB.Save(&invoice)

	return invoice
}
