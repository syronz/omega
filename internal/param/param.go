package param

import (
	"omega/internal/types"
	"omega/pkg/dict"
)

// Param for describing request's parameter
type Param struct {
	Pagination
	Search       string
	PreCondition string
	UserID       types.RowID
	Language     dict.Language
}

// Pagination is a struct, contains the fields which affected the front-end pagination
type Pagination struct {
	Select string
	Order  string
	Limit  uint64
	Offset uint64
}
