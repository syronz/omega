package basmodel

import (
	"omega/internal/core"
	"omega/internal/core/coract"
	"omega/internal/term"
	"omega/internal/types"
)

const (
	// RoleTable is used inside the repo layer
	RoleTable = "bas_roles"
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
func (p Role) Columns(variate string) (string, error) {
	full := []string{
		"bas_roles.id",
		"bas_roles.name",
		"bas_roles.description",
		"bas_roles.resources",
		"bas_roles.created_at",
		"bas_roles.updated_at",
	}

	return core.CheckColumns(full, variate)
}

// Validate check the type of fields
func (p *Role) Validate(act coract.Action) error {
	fieldError := core.NewFieldError(term.Error_in_role_form)

	switch act {
	case coract.Save:
		if p.Name == "" {
			fieldError.Add(term.V_is_required, "Name", "name")
		}

		if len(p.Name) < 5 {
			fieldError.Add(term.Name_at_least_be_5_character, nil, "name")
		}

		if p.Resources == "" {
			fieldError.Add(term.V_is_required, "Resources", "resources")
		}

	}

	if fieldError.HasError() {
		return fieldError
	}
	return nil

}
