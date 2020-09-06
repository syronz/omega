package basmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// Auth model
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate check the type of fields for auth
func (p *Auth) Validate(act coract.Action) error {
	var err error
	// fieldError := corerr.NewSilent("E1014832", params, base.Domain, nil)
	// err = limberr.AddCode(err, code)

	switch act {
	case coract.Login:
		if p.Username == "" {
			// fieldError.Add("username", corerr.V_is_required, "username")
			err = limberr.AddInvalidParam(err, "username", corerr.V_is_required, dict.R("username"))
		}

		if p.Password == "" {
			// fieldError.Add("password", corerr.V_is_required, "password")
			err = limberr.AddInvalidParam(err, "username", corerr.V_is_required, "username")
		}
	}

	// return fieldError.Final()
	return err
}
