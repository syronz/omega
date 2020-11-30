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
	AccountID types.RowID `json:"account_id"`
	PhoneID   types.RowID `json:"phone_id"`
	Default   byte        `gorm:"default:0" json:"default"`
}
