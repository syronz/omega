package basmodel

import (
	"omega/internal/core"
	"omega/internal/core/action"
	"omega/internal/term"
)

// BasAuth model
type BasAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Validate check the type of fields for auth
func (p *BasAuth) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_login_form)

	switch act {
	case action.Login:
		if p.Username == "" {
			fieldError.Add(term.V_is_required, "Username", "username")
		}

		if p.Password == "" {
			fieldError.Add(term.V_is_required, "Password", "password")
		}
	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil
}
