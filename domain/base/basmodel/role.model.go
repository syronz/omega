package basmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// RoleTable is a global instance for working with role
const (
	RoleTable = "bas_roles"
)

// Role model
type Role struct {
	types.GormCol
	Name        string `gorm:"not null;unique" json:"name,omitempty"`
	Resources   string `gorm:"type:text" json:"resources,omitempty"`
	Description string `json:"description,omitempty"`
}

// Validate check the type of fields
func (p *Role) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:

		if len(p.Name) < 5 {
			err = limberr.AddInvalidParam(err, "name",
				corerr.MinimumAcceptedCharacterForVisV,
				dict.R(corterm.Name), 5)
		}

		if len(p.Name) > 255 {
			err = limberr.AddInvalidParam(err, "name",
				corerr.MaximumAcceptedCharacterForVisV,
				dict.R(corterm.Name), 255)
		}

		if p.Resources == "" {
			err = limberr.AddInvalidParam(err, "resources",
				corerr.VisRequired, dict.R(corterm.Resources))
		}

		if len(p.Description) > 255 {
			err = limberr.AddInvalidParam(err, "description",
				corerr.MaximumAcceptedCharacterForVisV,
				dict.R(corterm.Description), 255)
		}
	}

	return err
}
