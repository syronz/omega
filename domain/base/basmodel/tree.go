package basmodel

import "omega/internal/types"

// Tree is used for creating chart of accoutns
type Tree struct {
	ID        types.RowID  `json:"id"`
	CompanyID uint64       `json:"company_id"`
	NodeID    uint64       `json:"node_id"`
	ParentID  *types.RowID `json:"parent_id"`
	Code      string       `json:"code"`
	Name      string       `json:"name"`
	Type      types.Enum   `json:"type"`
	Children  []*Tree      `json:"children"`
	Counter   int          `json:"counter"`
}
