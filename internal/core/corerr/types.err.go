package corerr

import (
	"net/http"
	"omega/domain/base"
	"omega/pkg/limberr"
)

var UnauthorizedErr limberr.CustomError = "unauthorized"
var NotFoundErr limberr.CustomError = "not found"
var ValidationFailedErr limberr.CustomError = "validation failed"

var UniqErrorMap limberr.CustomErrorMap

func init() {
	UniqErrorMap = make(map[limberr.CustomError]limberr.ErrorTheme)

	UniqErrorMap[UnauthorizedErr] = limberr.ErrorTheme{
		Type:   "#Unauthorized",
		Title:  Unauthorized,
		Domain: base.Domain,
		Status: http.StatusUnauthorized,
	}

	UniqErrorMap[ValidationFailedErr] = limberr.ErrorTheme{
		Type:   "#VALIDATION_FAILED",
		Title:  Validation_failed,
		Domain: base.Domain,
		Status: http.StatusUnprocessableEntity,
	}

	UniqErrorMap[NotFoundErr] = limberr.ErrorTheme{
		Type:  "#NOT_FOUND",
		Title: RecordNotFound,
	}
}

/*
// NotFound is used when findbyid returns nill
func (p CustomError) NotFound(part, field string, value interface{}, path string) error {
	field = dict.T(field, p.Lang)
	part = dict.T(part, p.Lang)
	return &CustomError{
		Code:          p.Code,
		Domain:        p.Domain,
		Type:          p.ErrPanel + string(p.Lang) + ".html#NOT_FOUND",
		Title:         dict.T(RecordNotFound, p.Lang),
		Message:       dict.T(Record__NotFoundIn_, p.Lang, field, value, part),
		MessageParams: []interface{}{field, fmt.Sprint(value), part},
		Path:          path,
		Status:        http.StatusNotFound,
		OriginalError: p.OriginalError,
	}
}

// NotBind is used when findbyid returns nill
func (p CustomError) NotBind(field string, path string) error {
	field = dict.T(field, p.Lang)
	return &CustomError{
		Code:          p.Code,
		Domain:        p.Domain,
		Type:          p.ErrPanel + string(p.Lang) + ".html#NOT_BIND",
		Title:         dict.T(Bind_failed, p.Lang),
		Message:       dict.T(V_is_not_valid, p.Lang, field),
		MessageParams: []interface{}{field},
		Path:          path,
		Status:        http.StatusUnprocessableEntity,
		OriginalError: p.OriginalError,
	}
}


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
*/
