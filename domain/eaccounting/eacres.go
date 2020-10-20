package eaccounting

import "omega/internal/types"

// list of resources for eaccounting domain
const (
	Domain string = "eaccounting"

	CurrencyWrite types.Resource = "currency:write"
	CurrencyRead  types.Resource = "currency:read"
	CurrencyExcel types.Resource = "currency:excel"

	TransactionRead  types.Resource = "transaction:read"
	TransactionWrite types.Resource = "transaction:write"
	TransactionExcel types.Resource = "transaction:excel"
)
