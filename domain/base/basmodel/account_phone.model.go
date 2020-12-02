package basmodel

import (
	"omega/internal/types"
)

// AccountPhoneTable is used inside the repo layer
const (
	AccountPhoneTable = "bas_account_phones"
)

// AccountPhone model
type AccountPhone struct {
	types.FixedNode
	AccountID types.RowID `gorm:"not null;unique_index:idx_acount_phone" json:"account_id"`
	PhoneID   types.RowID `gorm:"not null;unique_index:idx_acount_phone" json:"phone_id"`
	Default   byte        `gorm:"default:0" json:"default"`
}
