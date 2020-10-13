package accounttype

import (
	"omega/internal/types"
)

const (
	Eternal  types.Enum = "eternal"
	Asset    types.Enum = "asset"
	Expense  types.Enum = "expense"
	Trader   types.Enum = "trader"
	Provider types.Enum = "provider"
	Cashier  types.Enum = "cashier"
	Fee      types.Enum = "fee"
	Fixer    types.Enum = "fixer"
)

var List = []types.Enum{
	Eternal,
	Asset,
	Expense,
	Trader,
	Provider,
	Cashier,
	Fee,
	Fixer,
}

var ForbiddenNegative = []types.Enum{
	Trader,
	Provider,
	Cashier,
}

// Join make a string for showing in the api
func Join() string {
	return types.JoinEnum(List)
}
