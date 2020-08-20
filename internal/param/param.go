package param

import (
	"omega/internal/types"
)

// Param for describing request's parameter
type Param struct {
	Pagination
	Search       string
	PreCondition string
	UserID       types.RowID
	Language     string
}

// Pagination is a struct, contains the fields which affected the front-end pagination
type Pagination struct {
	Select string
	Order  string
	Limit  uint64
	Offset uint64
}
