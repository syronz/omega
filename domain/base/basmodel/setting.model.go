package basmodel

import (
	"omega/internal/core"
	"omega/internal/core/action"
	"omega/internal/term"
	"omega/internal/types"
	"omega/utils/helper"
)

const (
	// SettingTable is used inside the repo layer
	SettingTable = "bas_settings"
)

// Setting model
type Setting struct {
	types.FixedCol
	Property    types.Setting `gorm:"not null;unique_index:idx_companyID_property" json:"property,omitempty"`
	Value       string        `gorm:"type:text" json:"value,omitempty"`
	Type        string        `json:"type,omitempty"`
	Description string        `json:"description,omitempty"`
}

// Pattern returns the search pattern to be used inside the gorm's where
func (p Setting) Pattern() string {
	return `(
		bas_settings.property LIKE '%[1]v%%' OR
		bas_settings.ID = '%[1]v' OR
		bas_settings.value LIKE '%[1]v' OR
		bas_settings.type LIKE '%[1]v' OR
		bas_settings.description LIKE '%[1]v'
	)`
}

// Columns return list of total columns according to request, useful for inner joins
func (p Setting) Columns(variate string) (string, error) {
	full := []string{
		"bas_settings.id",
		"bas_settings.property",
		"bas_settings.value",
		"bas_settings.type",
		"bas_settings.description",
	}

	return helper.CheckColumns(full, variate)
}

// Validate check the type of fields
func (p *Setting) Validate(act action.Action) error {
	fieldError := core.NewFieldError(term.Error_in_role_form)

	switch act {
	case action.Save:
		if p.Property == "" {
			fieldError.Add(term.V_is_required, "Property", "property")
		}
	}

	if fieldError.HasError() {
		return fieldError
	}

	return nil

}
