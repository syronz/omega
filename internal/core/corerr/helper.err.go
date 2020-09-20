package corerr

import (
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// RecordNotFoundHelper reduce the code size in case of NotFoundErr happened
func RecordNotFoundHelper(err error, code string, v1 string, v2 interface{}, v3 string) error {
	return limberr.Take(err, code).
		Message(RecordVVNotFoundInV, dict.R(v1), v2, dict.R(v3)).
		Custom(NotFoundErr).Build()
}

// InternalServerErrorHelper help generate proper error in case of internal server error happened
func InternalServerErrorHelper(err error, code string) error {
	return limberr.Take(err, code).
		Message(InternalServerError).
		Custom(InternalServerErr).Build()
}

// ValidationFailedHelper help generate proper error in case of validation faced error
func ValidationFailedHelper(err error, code string) error {
	return limberr.Take(err, code).
		Message(ValidationFailed).
		Custom(ValidationFailedErr).Build()
}
