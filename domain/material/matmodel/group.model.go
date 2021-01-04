package matmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// GroupTable is a global instance for working with group
const (
	GroupTable = "mat_groups"
)

// Group model
type Group struct {
	types.FixedCol
	ParentID  *types.RowID `json:"parent_id,omitempty"`
	Name      string       `gorm:"not null;unique" json:"name,omitempty"`
	LatinName string       `gorm:"not null;unique" json:"latin_name,omitempty"`
	Code      string       `gorm:"not null;unique" json:"code,omitempty"`
	Leaf      bool         `json:"leaf,omitempty"`
	Notes     string       `json:"notes,omitempty"`
	Caption   string       `json:"caption,omitempty"`
}

// Validate check the type of fields
func (p *Group) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:

		if p.Code == "" {
			err = limberr.AddInvalidParam(err, "code",
				corerr.VisRequired,
				dict.R(corterm.Code))
		}

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

	}

	return err
}
