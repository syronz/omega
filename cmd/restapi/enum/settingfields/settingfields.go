package settingfields

import (
	"omega/internal/types"
	"strings"
)

const (
	CompanyName          types.Setting = "company_name"
	ReceiptHeader        types.Setting = "receipt_header"
	DefaultLang          types.Setting = "default_language"
	CompanyLogo          types.Setting = "company_logo"
	InvoiceLogo          types.Setting = "invoice_logo"
	InvoiceNumberPattern types.Setting = "invoice_number_pattern"
	ReceiptPhone         types.Setting = "receipt_phone"
	ReceiptAddress       types.Setting = "receipt_address"
)

var List = []types.Setting{
	CompanyName,
	ReceiptHeader,
	DefaultLang,
	CompanyLogo,
	InvoiceLogo,
	InvoiceNumberPattern,
	ReceiptPhone,
	ReceiptAddress,
}

// Join make a string for showing in the api
func Join() string {
	var strArr []string

	for _, v := range List {
		strArr = append(strArr, string(v))
	}

	return strings.Join(strArr, ", ")
}
