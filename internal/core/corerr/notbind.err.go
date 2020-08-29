package corerr

import (
	"net/http"
	"omega/pkg/dict"
)

// NotBind is used when findbyid returns nill
func (p CustomError) NotBind(field string, path string) error {
	field = dict.T(field, p.Lang)
	return &CustomError{
		Code:          p.Code,
		Domain:        p.Domain,
		Type:          p.ErrPanel + string(p.Lang) + ".html#NOT_BIND",
		Title:         dict.T(Bind_failed, p.Lang),
		Message:       dict.T(_V_is_not_valid, p.Lang, field),
		MessageParams: []string{field},
		Path:          path,
		Status:        http.StatusUnprocessableEntity,
		OriginalError: p.OriginalError,
	}
}
