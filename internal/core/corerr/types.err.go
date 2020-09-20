package corerr

import (
	"net/http"
	"omega/domain/base"
	"omega/pkg/limberr"
)

const (
	Nil limberr.CustomError = iota
	UnkownErr
	UnauthorizedErr
	NotFoundErr
	RouteNotFoundErr
	ValidationFailedErr
	ForeignErr
	DuplicateErr
	InternalServerErr
	BindingErr
)

// UniqErrorMap is used for categorized errors and connect error with error page also primary fill
// the status code and domain and title
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
		Title:  ValidationFailed,
		Domain: base.Domain,
		Status: http.StatusUnprocessableEntity,
	}

	UniqErrorMap[NotFoundErr] = limberr.ErrorTheme{
		Type:   "#NOT_FOUND",
		Title:  RecordNotFound,
		Domain: base.Domain,
		Status: http.StatusNotFound,
	}

	UniqErrorMap[RouteNotFoundErr] = limberr.ErrorTheme{
		Type:   "#NOT_FOUND",
		Title:  RouteNotFound,
		Domain: base.Domain,
		Status: http.StatusNotFound,
	}

	UniqErrorMap[ForeignErr] = limberr.ErrorTheme{
		Type:   "#FOREIGN_KEY",
		Title:  ErrorBecauseOfForeignKey,
		Domain: base.Domain,
		Status: http.StatusConflict,
	}

	UniqErrorMap[InternalServerErr] = limberr.ErrorTheme{
		Type:   "#INTERNAL_SERVER_ERROR",
		Title:  InternalServerError,
		Domain: base.Domain,
		Status: http.StatusInternalServerError,
	}

	UniqErrorMap[DuplicateErr] = limberr.ErrorTheme{
		Type:   "#DUPLICATE_ERROR",
		Title:  DuplicateHappened,
		Domain: base.Domain,
		Status: http.StatusConflict,
	}

	UniqErrorMap[BindingErr] = limberr.ErrorTheme{
		Type:   "#NOT_BIND",
		Title:  BindFailed,
		Domain: base.Domain,
		Status: http.StatusUnprocessableEntity,
	}
}
