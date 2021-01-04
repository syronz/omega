package synmodel

import (
	"omega/domain/sync/enum/companytype"
	"omega/domain/sync/synterm"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"time"
)

// CompanyTable is used inside the repo layer
const (
	CompanyTable = "syn_companies"
)

// Company model
type Company struct {
	types.GormCol
	Name          string     `gorm:"not null" json:"name,omitempty"`
	LegalName     string     `gorm:"not null;unique" json:"legal_name,omitempty"`
	Key           string     `gorm:"type:text" json:"key,omitempty"`
	ServerAddress string     `json:"server_address,omitempty"`
	Expiration    *time.Time `json:"expiration,omitempty"`
	License       string     `gorm:"unique" json:"license,omitempty"`
	Plan          string     `json:"plan,omitempty"`
	Detail        string     `json:"detail,omitempty"`
	Phone         string     `gorm:"not null" json:"phone,omitempty"`
	Email         string     `gorm:"not null" json:"email,omitempty"`
	Website       string     `gorm:"not null" json:"website,omitempty"`
	Type          string     `gorm:"not null" json:"type,omitempty"`
	Code          string     `gorm:"not null" json:"code,omitempty"`
	Logo          string     `json:"logo,omitempty"`
	Banner        string     `json:"banner,omitempty"`
	Footer        string     `json:"footer,omitempty"`
	AdminUsername string     `gorm:"-" json:"admin_username,omitempty" table:"-"`
	AdminPassword string     `gorm:"-" json:"admin_password,omitempty" table:"-"`
	Lang          dict.Lang  `gorm:"-" josn:"lang,omitempty" table:"-"`
}

// Validate check the type of
func (p *Company) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Save:

		if len(p.Name) < 1 {
			err = limberr.AddInvalidParam(err, "name",
				corerr.MinimumAcceptedCharacterForVisV,
				dict.R(corterm.Name), 1)
		}

		if len(p.Name) > 255 {
			err = limberr.AddInvalidParam(err, "name",
				corerr.MaximumAcceptedCharacterForVisV,
				dict.R(corterm.Name), 255)
		}

		if p.Code == "" {
			err = limberr.AddInvalidParam(err, "code",
				corerr.VisRequired, dict.R(corterm.Code))
		}

		if len(p.Detail) > 255 {
			err = limberr.AddInvalidParam(err, "detail",
				corerr.MaximumAcceptedCharacterForVisV,
				dict.R(synterm.Detail), 255)
		}

		if ok, _ := helper.Includes(companytype.List, types.Enum(p.Type)); !ok {
			return limberr.AddInvalidParam(err, "type",
				corerr.AcceptedValueForVareV, dict.R(corterm.Type),
				companytype.Join())
		}
	}

	return err

}
