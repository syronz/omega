package companytype

// Company types
const (
	Base                        = "base"
	MultiBranchCentralFinance   = "multi branch with centeral finance"
	MultiBranchScatteredFinance = "multi branch with scattered finance"
	CarpetStore                 = "carpet store"
	SimplePOS                   = "simple POS"
	Other                       = "other"
)

// Types is list of account types for checking
var Types = []string{
	Base,
	MultiBranchCentralFinance,
	MultiBranchScatteredFinance,
	CarpetStore,
	SimplePOS,
	Other,
}
