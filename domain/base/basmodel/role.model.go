package basmodel

import (
	"omega/internal/core/coract"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
)

const (
	// RoleTable is used inside the repo layer
	RoleTable = "bas_roles"
	RolePart  = "role"
	RolesPart = "roles"
)

// Role model
type Role struct {
	types.GormCol
	Name        string `gorm:"not null;unique" json:"name,omitempty"`
	Resources   string `gorm:"type:text" json:"resources,omitempty"`
	Description string `json:"description,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Role) Pattern() string {
	return `(bas_roles.name LIKE '%[1]v%%' OR
		bas_roles.id = '%[1]v' OR
		bas_roles.description LIKE '%%%[1]v%%' OR
		bas_roles.resources LIKE '%%%[1]v%%')`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Role) Columns(variate string, params param.Param) (string, error) {
	full := []string{
		"bas_roles.id",
		"bas_roles.name",
		"bas_roles.description",
		"bas_roles.resources",
		"bas_roles.created_at",
		"bas_roles.updated_at",
	}

	return validator.CheckColumns(full, variate, params)
}

// Validate check the type of fields
func (p *Role) Validate(act coract.Action, params param.Param) error {
	// fieldError := corerr.NewSilent("E1062183", params, base.Domain, nil).
	// 	FieldError("roles/[:roleID]", corerr.Validation_failed_for_V, dict.R("role"))

	// switch act {
	// case coract.Save:
	// 	if p.Name == "" {
	// 		fieldError.Add("name", corerr.V_is_required, dict.R("Name"))
	// 	}

	// 	if len(p.Name) < 5 {
	// 		fieldError.Add("name", corerr.Minimum_accepted_character_for_V_is_V,
	// 			dict.R("Name"), 5)
	// 	}

	// 	if p.Resources == "" {
	// 		fieldError.Add("resources", corerr.V_is_required, "Resources")
	// 	}

	// }

	// return fieldError.Final()
	return nil
}
