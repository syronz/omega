package invoice

import (
	"omega/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

func ProvideInvoiceRepostiory(c config.CFG) InvoiceRepository {
	return InvoiceRepository{DB: c.DB}
}

func (p *InvoiceRepository) Save(invoice Invoice) Invoice {
	p.DB.Save(&invoice)

	return invoice
}
