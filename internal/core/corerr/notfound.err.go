package corerr

import (
	"fmt"
	"net/http"
	"omega/pkg/dict"
)

// NotFound is used when findbyid returns nill
func (p CustomError) NotFound(part, field string, value interface{}, path string) error {
	field = dict.T(field, p.Lang)
	part = dict.T(part, p.Lang)
	return &CustomError{
		Code:          p.Code,
		Type:          p.ErrPanel,
		Title:         dict.T(RecordNotFound, p.Lang),
		Message:       dict.T(Record__NotFoundIn_, p.Lang, field, value, part),
		MessageParams: []string{field, fmt.Sprint(value), part},
		Path:          path,
		Status:        http.StatusNotFound,
		OriginalError: p.OriginalError,
	}
}
