package companytype

import "omega/internal/types"

// Company types
const (
	Base                        types.Enum = "base"
	MultiBranchCentralFinance   types.Enum = "multi branch with centeral finance"
	MultiBranchScatteredFinance types.Enum = "multi branch with scattered finance"
	SimplePOS                   types.Enum = "simple POS"
	Other                       types.Enum = "other"
)

var List = []types.Enum{
	Base,
	MultiBranchCentralFinance,
	MultiBranchScatteredFinance,
	SimplePOS,
	Other,
}

// Join make a string for showing in the api
func Join() string {
	return types.JoinEnum(List)
}
