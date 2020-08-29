package corerr

import (
	"net/http"
	"omega/pkg/dict"
)

// InternalServer is used when findbyid returns nill
func (p CustomError) InternalServer(part, field string, value interface{}, path string) error {
	field = dict.T(field, p.Lang)
	part = dict.T(part, p.Lang)
	return &CustomError{
		Code:          p.Code,
		Type:          p.ErrPanel + string(p.Lang),
		Title:         dict.T(InternalServerError, p.Lang),
		Message:       dict.T(Internal_Server_Error_Happened_Please_Aware_Administration_And_Gave_Him_Error_Code, p.Lang),
		Path:          path,
		Status:        http.StatusInternalServerError,
		OriginalError: p.OriginalError,
	}
}
