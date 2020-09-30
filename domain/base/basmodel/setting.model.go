package basmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// SettingTable is used inside the repo layer for specify the table name
const (
	SettingTable = "bas_settings"
)

// Setting model
type Setting struct {
	types.GormCol
	Property    types.Setting `gorm:"not null;unique_index:idx_companyID_property" json:"property,omitempty"`
	Value       string        `gorm:"type:text" json:"value,omitempty"`
	Type        string        `json:"type,omitempty"`
	Description string        `json:"description,omitempty"`
}

// Validate check the type of fields
func (p *Setting) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:
		if p.Property == "" {
			err = limberr.AddInvalidParam(err, "property",
				corerr.VisRequired, dict.R(corterm.Property))
		}
		fallthrough
	case coract.Update:
		if p.Value == "" {
			err = limberr.AddInvalidParam(err, "value",
				corerr.VisRequired, dict.R(corterm.Value))
		}
	}

	return err
}
