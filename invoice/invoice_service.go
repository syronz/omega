package invoice

type InvoiceService struct {
	InvoiceRepository InvoiceRepository
}

func ProvideInvoiceService(p InvoiceRepository) InvoiceService {
	return InvoiceService{InvoiceRepository: p}
}

func (p *InvoiceService) Save(invoice Invoice) Invoice {
	p.InvoiceRepository.Save(invoice)

	return invoice
}
