package invoice

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type InvoiceAPI struct {
	InvoiceService InvoiceService
}

func ProvideInvoiceAPI(p InvoiceService) InvoiceAPI {
	return InvoiceAPI{InvoiceService: p}
}



func (p *InvoiceAPI) Create(c *gin.Context) {
	var invoiceDTO InvoiceDTO
	err := c.BindJSON(&invoiceDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdInvoice := p.InvoiceService.Save(ToInvoice(invoiceDTO))

	c.JSON(http.StatusOK, gin.H{"invoice": ToInvoiceDTO(createdInvoice)})
}

