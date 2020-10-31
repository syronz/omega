package eaccounting

import "omega/internal/types"

// list of resources for eaccounting domain
const (
	Domain string = "eaccounting"

	CurrencyWrite types.Resource = "currency:write"
	CurrencyRead  types.Resource = "currency:read"
	CurrencyExcel types.Resource = "currency:excel"

	TransactionRead   types.Resource = "transaction:read"
	TransactionManual types.Resource = "transaction:manual"
	TransactionUpdate types.Resource = "transaction:update"
	TransactionDelete types.Resource = "transaction:delete"
	TransactionExcel  types.Resource = "transaction:excel"
)
