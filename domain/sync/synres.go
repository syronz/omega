package sync

import "omega/internal/types"

// list of resources for sync domain
const (
	Domain string = "sync"

	CompanyRead   types.Resource = "sync-company:read"
	CompanyManual types.Resource = "sync-company:manual"
	CompanyUpdate types.Resource = "sync-company:update"
	CompanyDelete types.Resource = "sync-company:delete"
	CompanyExcel  types.Resource = "sync-company:excel"
)
