package basmodel

import (
	"omega/domain/base"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/param"
)

// Auth model
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate check the type of fields for auth
func (p *Auth) Validate(act coract.Action, params param.Param) error {
	fieldError := corerr.NewSilent("E1014832", params, base.Domain, nil)

	switch act {
	case coract.Login:
		if p.Username == "" {
			fieldError.Add("username", corerr.V_is_required, "username")
		}

		if p.Password == "" {
			fieldError.Add("password", corerr.V_is_required, "password")
		}
	}

	return fieldError.Final()
}
