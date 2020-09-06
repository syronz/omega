package basmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// Auth model
type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate check the type of fields for auth
func (p *Auth) Validate(act coract.Action) (err error) {

	switch act {
	case coract.Login:
		if p.Username == "" {
			err = limberr.AddInvalidParam(err, "username",
				corerr.V_is_required, dict.R(corterm.Username))
		}

		if p.Password == "" {
			err = limberr.AddInvalidParam(err, "password",
				corerr.V_is_required, dict.R(corterm.Password))
		}
	}

	return err
}
