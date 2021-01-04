package eacmodel

import (
	"omega/domain/eaccounting/enum/transactiontype"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"time"
)

// TransactionTable is a global instance for working with transaction
const (
	TransactionTable = "eac_transactions"
)

// Transaction model
type Transaction struct {
	types.FixedNode
	CurrencyID  types.RowID `gorm:"not null" json:"currency_id,omitempty"`
	CreatedBy   types.RowID `gorm:"not null" json:"created_by,omitempty"`
	Hash        string      `gorm:"not null;unique" json:"hash,omitempty"`
	Type        types.Enum  `json:"type,omitempty"`
	Description *string     `json:"description,omitempty"`
	Amount      float64     `json:"amount,omitempty"`
	Pioneer     types.RowID `gorm:"-" json:"pioneer,omitempty" table:"-"`
	Follower    types.RowID `gorm:"-" json:"follower,omitempty" table:"-"`
	PostDate    time.Time   `gorm:"-" json:"post_date,omitempty" table:"-"`
	Slots       []Slot      `gorm:"-" json:"slots,omitempty" table:"-"`
}

// Validate check the type of fields
func (p *Transaction) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:
		if p.Description != nil {
			if len(*p.Description) > 255 {
				err = limberr.AddInvalidParam(err, "description",
					corerr.MaximumAcceptedCharacterForVisV,
					dict.R(corterm.Description), 255)
			}
		}

		if ok, _ := helper.Includes(transactiontype.List, p.Type); !ok {
			return limberr.AddInvalidParam(err, "type",
				corerr.AcceptedValueForVareV, dict.R(corterm.Type),
				transactiontype.Join())
		}
	}

	return err
}
