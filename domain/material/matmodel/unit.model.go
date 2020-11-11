package matmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// UnitTable is a global instance for working with unit
const (
	UnitTable = "mat_units"
)

// Unit model
type Unit struct {
	types.FixedCol
	Name        string  `gorm:"not null;unique" json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Validate check the type of fields
func (p *Unit) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:

		if len(p.Name) < 2 {
			err = limberr.AddInvalidParam(err, "name",
				corerr.MinimumAcceptedCharacterForVisV,
				dict.R(corterm.Name), 2)
		}

		if len(p.Name) > 255 {
			err = limberr.AddInvalidParam(err, "name",
				corerr.MaximumAcceptedCharacterForVisV,
				dict.R(corterm.Name), 255)
		}

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
