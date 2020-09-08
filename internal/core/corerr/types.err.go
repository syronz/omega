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
		Type:   "#NOT_FOUND",
		Title:  RecordNotFound,
		Domain: base.Domain,
		Status: http.StatusNotFound,
	}
}

/*

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
