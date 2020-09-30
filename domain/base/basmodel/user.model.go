package basmodel

import (
	"omega/domain/base/message/basterm"
	"omega/internal/consts"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
	"omega/pkg/dict"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"regexp"
	"strings"
)

// UserTable is used inside the repo layer
const (
	UserTable = "bas_users"
)

// User model
type User struct {
	ID        types.RowID `gorm:"not null;unique" json:"id"`
	RoleID    types.RowID `gorm:"index:role_id_idx" json:"role_id"`
	Username  string      `gorm:"not null;unique" json:"username,omitempty"`
	Password  string      `gorm:"not null" json:"password,omitempty"`
	Lang      dict.Lang   `gorm:"type:varchar(2);default:'en'" json:"lang,omitempty"`
	Email     string      `json:"email,omitempty"`
	Extra     interface{} `sql:"-" json:"user_extra,omitempty" table:"-"`
	Resources string      `sql:"-" json:"resources,omitempty" table:"bas_roles.resources"`
	Role      string      `sql:"-" json:"role,omitempty" table:"bas_roles.name as role"`
}

// Validate check the type of
func (p *User) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Create:

		err = validateUserUsername(err, p.Username)
		err = validateUserPassword(err, p.Password)
		err = validateUserRole(err, p.RoleID)
		err = validateUserLang(err, p.Lang)
		err = validateUserEmail(err, p.Email)

	case coract.Update:

		err = validateUserUsername(err, p.Username)

		if p.Password != "" {
			err = validateUserPassword(err, p.Password)
		}

		err = validateUserRole(err, p.RoleID)
		err = validateUserLang(err, p.Lang)
		err = validateUserEmail(err, p.Email)

	//for default we validate all fields
	default:
		err = validateUserUsername(err, p.Username)
		err = validateUserPassword(err, p.Password)
		err = validateUserRole(err, p.RoleID)
		err = validateUserLang(err, p.Lang)
		err = validateUserEmail(err, p.Email)
	}

	return err
}

func validateUserPassword(err error, password string) error {
	if len(password) < consts.MinimumPasswordChar {
		return limberr.AddInvalidParam(err, "password",
			corerr.MinimumAcceptedCharacterForVisV,
			dict.R(corterm.Password), consts.MinimumPasswordChar)
	}
	return err
}

func validateUserUsername(err error, username string) error {
	if username == "" {
		return limberr.AddInvalidParam(err, "username",
			corerr.VisRequired, dict.R(corterm.Username))
	}
	return err
}

func validateUserRole(err error, roleID types.RowID) error {
	if roleID == 0 {
		return limberr.AddInvalidParam(err, "role_id",
			corerr.VisRequired, dict.R(basterm.Role))
	}
	return err
}

func validateUserLang(err error, lang dict.Lang) error {
	if ok, _ := helper.Includes(dict.Langs, lang); !ok {
		var str []string
		for _, v := range dict.Langs {
			str = append(str, string(v))
		}
		return limberr.AddInvalidParam(err, "language",
			corerr.AcceptedValueForVareV, dict.R(corterm.Language),
			strings.Join(str, ", "))
	}
	return err
}

func validateUserEmail(err error, email string) error {
	if email != "" {
		re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !re.MatchString(email) {
			return limberr.AddInvalidParam(err, "email",
				corerr.VisNotValid, dict.R(corterm.Email))
		}
	}
	return err
}
