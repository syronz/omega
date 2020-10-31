package param

import (
	"omega/internal/consts"
	"omega/internal/types"
	"omega/pkg/dict"
)

// Param for describing request's parameter
type Param struct {
	Pagination
	Search       string
	Filter       string
	PreCondition string
	UserID       types.RowID
	CompanyID    uint64
	NodeID       uint64
	Lang         dict.Lang
	ErrPanel     string
}

// Pagination is a struct, contains the fields which affected the front-end pagination
type Pagination struct {
	Select string
	Order  string
	Limit  uint64
	Offset uint64
}

func New() Param {
	var param Param
	param.Limit = consts.DefaultLimit

	return param
}
