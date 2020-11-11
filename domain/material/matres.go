package material

import "omega/internal/types"

// list of resources for material domain
const (
	Domain string = "material"

	CompanyWrite types.Resource = "company:write"
	CompanyRead  types.Resource = "company:read"
	CompanyExcel types.Resource = "company:excel"

	ColorWrite types.Resource = "color:write"
	ColorRead  types.Resource = "color:read"
	ColorExcel types.Resource = "color:excel"

	GroupWrite types.Resource = "group:write"
	GroupRead  types.Resource = "group:read"
	GroupExcel types.Resource = "group:excel"

	UnitWrite types.Resource = "unit:write"
	UnitRead  types.Resource = "unit:read"
	UnitExcel types.Resource = "unit:excel"
)
