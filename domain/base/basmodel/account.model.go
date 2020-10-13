package basmodel

import (
	"omega/domain/base/enum/accounttype"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
)

// AccountTable is used inside the repo layer
const (
	AccountTable = "bas_accounts"
)

// Account model
type Account struct {
	types.GormCol
	Name   string     `gorm:"not null;unique" json:"name,omitempty"`
	Type   types.Enum `json:"type,omitempty"`
	Status types.Enum `json:"status,omitempty"`
}

// Validate check the type of fields
func (p *Account) Validate(act coract.Action) (err error) {

	// switch act {
	// case coract.Save:

	// 	if len(p.Name) < 5 {
	// 		err = limberr.AddInvalidParam(err, "name",
	// 			corerr.MinimumAcceptedCharacterForVisV,
	// 			dict.R(corterm.Name), 5)
	// 	}

	// 	if len(p.Name) > 255 {
	// 		err = limberr.AddInvalidParam(err, "name",
	// 			corerr.MaximumAcceptedCharacterForVisV,
	// 			dict.R(corterm.Name), 255)
	// 	}

	// 	if p.Resources == "" {
	// 		err = limberr.AddInvalidParam(err, "resources",
	// 			corerr.VisRequired, dict.R(corterm.Resources))
	// 	}

	// 	if len(p.Description) > 255 {
	// 		err = limberr.AddInvalidParam(err, "description",
	// 			corerr.MaximumAcceptedCharacterForVisV,
	// 			dict.R(corterm.Description), 255)
	// 	}
	// }

	// TODO: it should be checked after API has been created
	if ok, _ := helper.Includes(accounttype.List, p.Type); !ok {
		// var str []string
		// for _, v := range dict.Langs {
		// 	str = append(str, string(v))
		// }
		return limberr.AddInvalidParam(err, "type",
			corerr.AcceptedValueForVareV, dict.R(corterm.Type),
			accounttype.Join())
	}

	return err
}
