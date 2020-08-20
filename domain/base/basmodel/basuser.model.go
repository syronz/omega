package basmodel

import (
	"omega/internal/core"
	"omega/internal/core/action"
	"omega/internal/core/lang"
	"omega/internal/term"
	"omega/internal/types"
	"omega/utils/helper"
	"regexp"
	"strings"
)

// BasUser model
type BasUser struct {
	ID        types.RowID `gorm:"not null;unique" json:"id"`
	RoleID    types.RowID `gorm:"index:role_id_idx" json:"role_id"`
	Username  string      `gorm:"not null;unique" json:"username,omitempty"`
	Password  string      `gorm:"not null" json:"password,omitempty"`
	Language  string      `gorm:"type:varchar(2);default:'en'" json:"language,omitempty"`
	Email     string      `json:"email,omitempty"`
	Extra     interface{} `sql:"-" json:"user_extra,omitempty"`
	Resources string      `sql:"-" json:"resources,omitempty"`
	Role      string      `sql:"-" json:"role,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p BasUser) Pattern() string {
	return `(
		bas_users.username LIKE '%[1]v%%' OR
		bas_users.id = '%[1]v' OR
		bas_users.email LIKE '%[1]v%%' OR
		bas_roles.name LIKE '%[1]v' 
	)`
}

// Columns return list of total columns according to request, useful for inner joins
func (p BasUser) Columns(variate string) (string, error) {
	full := []string{
		"bas_users.id",
		"bas_users.role_id",
		"bas_users.username",
		"bas_users.language",
		"bas_users.email",
		"bas_roles.name as role",
	}

	return helper.CheckColumns(full, variate)
}

// Validate check the type of
func (p *BasUser) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_users_form)

	switch act {
	case action.Create:

		if len(p.Password) < 8 {
			params := []interface{}{"password", 7}
			fieldError.Add(term.V_should_be_more_than_V_character, params, "password")
		}

		fallthrough

	case action.Update:

		if p.Username == "" {
			fieldError.Add(term.V_is_required, "Username", "username")
		}

		if p.RoleID == 0 {
			fieldError.Add(term.V_is_required, "Role", "role_id")
		}

		if ok, _ := helper.Includes(lang.Languages, p.Language); !ok {
			fieldError.Add(term.Accepted_values_are_v,
				strings.Join(lang.Languages, ", "), "language")
		}

		if p.Email != "" {
			re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
			if !re.MatchString(p.Email) {
				fieldError.Add(term.Email_is_not_valid, nil, "email")
			}
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil

}
