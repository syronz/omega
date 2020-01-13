package invoice

func ToInvoice(invoiceDTO InvoiceDTO) Invoice {
	return Invoice{InvoiceNumber: invoiceDTO.InvoiceNumber}
}

func ToInvoiceDTO(invoice Invoice) InvoiceDTO {
	return InvoiceDTO{ID: invoice.ID, InvoiceNumber: invoice.InvoiceNumber}
}

func ToInvoiceDTOs(invoices []Invoice) []InvoiceDTO {
	invoicedtos := make([]InvoiceDTO, len(invoices))

	for i, itm := range invoices {
		invoicedtos[i] = ToInvoiceDTO(itm)
	}

	return invoicedtos
}
