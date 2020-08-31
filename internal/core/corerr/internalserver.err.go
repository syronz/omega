package corerr

import (
	"net/http"
	"omega/pkg/dict"
)

// InternalServer is used when findbyid returns nill
func (p CustomError) InternalServer(path string) error {
	return &CustomError{
		Code:          p.Code,
		Domain:        p.Domain,
		Type:          p.ErrPanel + string(p.Lang) + ".html#INTERNAL_SERVER_ERROR",
		Title:         dict.T(InternalServerError, p.Lang),
		Message:       dict.T(Internal_Server_Error_Happened___, p.Lang),
		Path:          path,
		Status:        http.StatusInternalServerError,
		OriginalError: p.OriginalError,
	}
}
