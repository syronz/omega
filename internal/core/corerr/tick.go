package corerr

import (
	"omega/pkg/glog"
	"omega/pkg/limberr"
)

// Tick is combining AddCode and LogError in services to reduce the code
func Tick(err error, code string, message string, data ...interface{}) error {
	if code != "" {
		err = limberr.AddCode(err, code)
	}
	glog.LogError(err, message, data...)
	return err
}

// TickCustom is combining AddCode and LogError in services to reduce the code and specify the
// error type
func TickCustom(err error, custom limberr.CustomError, code string, message string,
	data ...interface{}) error {
	err = limberr.SetCustom(err, custom)
	return Tick(err, code, message, data...)
}

// TickValidate is automatically add validation error custom to the error
func TickValidate(err error, code string, message string,
	data ...interface{}) error {
	return TickCustom(err, ValidationFailedErr, code, message, data...)
}
