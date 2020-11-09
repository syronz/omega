package material

import "omega/internal/types"

// types for material domain
const (
	CreateCompany types.Event = "company-create"
	UpdateCompany types.Event = "company-update"
	DeleteCompany types.Event = "company-delete"
	ListCompany   types.Event = "company-list"
	ViewCompany   types.Event = "company-view"
	ExcelCompany  types.Event = "company-excel"
)
