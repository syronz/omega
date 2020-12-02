package sync

import "omega/internal/types"

// types for eaccounting domain
const (
	CreateCompany types.Event = "syn-company-create"
	UpdateCompany types.Event = "syn-company-update"
	DeleteCompany types.Event = "syn-company-delete"
	ListCompany   types.Event = "syn-company-list"
	ViewCompany   types.Event = "syn-company-view"
	ExcelCompany  types.Event = "syn-company-excel"
)
