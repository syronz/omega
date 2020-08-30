package basmodel

import (
	"omega/domain/base"
	"omega/internal/core/coract"
	"omega/internal/core/corerr"
	"omega/internal/core/validator"
	"omega/internal/param"
	"omega/internal/types"
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
func (p Setting) Columns(variate string, params param.Param) (string, error) {
	full := []string{
		"bas_settings.id",
		"bas_settings.property",
		"bas_settings.value",
		"bas_settings.type",
		"bas_settings.description",
	}

	return validator.CheckColumns(full, variate, params)
}

// Validate check the type of fields
func (p *Setting) Validate(act coract.Action, params param.Param) error {
	fieldError := corerr.NewSilent("E1072908", params, base.Domain, nil)

	switch act {
	case coract.Save:
		if p.Property == "" {
			fieldError.Add("property", corerr.V_is_required, "property")
		}
		fallthrough
	case coract.Update:
		if p.Value == "" {
			fieldError.Add("value", corerr.V_is_required, "value")
		}
	}

	return fieldError.Final()

}
