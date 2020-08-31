package corerr

import (
	"net/http"
	"omega/pkg/dict"
)

// FieldError is used when findbyid returns nill
func (p CustomError) FieldError(path string, msg string, msgParams ...interface{}) *CustomError {
	return &CustomError{
		Code:          p.Code,
		Domain:        p.Domain,
		Type:          p.ErrPanel + string(p.Lang) + ".html#VALIDATION_FAILED",
		Title:         dict.T(Validation_failed, p.Lang),
		Message:       dict.T(msg, p.Lang, msgParams...),
		MessageParams: msgParams,
		Path:          path,
		Status:        http.StatusUnprocessableEntity,
		OriginalError: p.OriginalError,
		Lang:          p.Lang,
	}
}

// SetMsg will update message and params
func (p *CustomError) SetMsg(msg string, msgParams ...interface{}) *CustomError {
	p.MessageParams = make([]interface{}, len(msgParams), len(msgParams))
	copy(p.MessageParams, msgParams)
	p.Message = dict.T(msg, p.Lang, msgParams...)
	return p
}

// Path will update path
func (p *CustomError) SetPath(path string) *CustomError {
	p.Path = path
	return p
}

// Add is used for add new element to the array of fields error
func (p *CustomError) Add(fieldName string, msg string, reasonParams ...interface{}) *CustomError {
	var field Field
	field.ReasonParams = make([]interface{}, len(reasonParams), len(reasonParams))
	copy(field.ReasonParams, reasonParams)
	field.Field = fieldName
	field.Reason = dict.T(msg, p.Lang, reasonParams...)
	p.InvalidParams = append(p.InvalidParams, field)
	return p
}

// Final helps to find out if param error exist or not
func (p *CustomError) Final() error {
	if len(p.InvalidParams) > 0 {
		return p
	}
	return nil
}
