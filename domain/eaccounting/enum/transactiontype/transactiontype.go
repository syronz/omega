package transactiontype

import "omega/internal/types"

// Transactions type
const (
	Manual types.Enum = "manual"
)

var List = []types.Enum{
	Manual,
}

// Join make a string for showing in the api
func Join() string {
	return types.JoinEnum(List)
}
