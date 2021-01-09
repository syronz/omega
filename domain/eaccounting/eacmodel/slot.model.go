package eacmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/limberr"
	"time"
)

// SlotTable is a global instance for working with slot
const (
	SlotTable = "eac_slots"
)

// Slot model
type Slot struct {
	types.FixedNode
	CurrencyID    types.RowID `gorm:"not null" json:"currency_id,omitempty"`
	AccountID     types.RowID `gorm:"not null" json:"account_id,omitempty"`
	TransactionID types.RowID `gorm:"not null" json:"transaction_id,omitempty"`
	Debit         float64     `json:"debit,omitempty"`
	Credit        float64     `json:"credit,omitempty"`
	Balance       float64     `json:"balance,omitempty"`
	Description   *string     `json:"description,omitempty"`
	Rows          int         `json:"rows,omitempty"`
	PostDate      time.Time   `json:"post_date,omitempty"`
}

// 	debit	credit	balance	description	post_date	row

// Validate check the type of fields
func (p *Slot) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:
		if p.Description != nil {
			if len(*p.Description) > 255 {
				err = limberr.AddInvalidParam(err, "description",
					corerr.MaximumAcceptedCharacterForVisV,
					dict.R(corterm.Description), 255)
			}
		}
	}

	return err
}
