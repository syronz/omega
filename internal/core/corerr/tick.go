package corerr

import (
	"omega/pkg/glog"
	"omega/pkg/limberr"
)

// Tick is combining AddCode and LogError in services to reduce the code
func Tick(err error, code string, message string, params ...interface{}) error {
	if code != "" {
		err = limberr.AddCode(err, code)
	}
	glog.LogError(err, message, params...)
	return err
}

// TickCustom is combining AddCode and LogError in services to reduce the code and specify the
// error type
func TickCustom(err error, custom limberr.CustomError, code string, message string,
	params ...interface{}) error {
	if code != "" {
		err = limberr.AddCode(err, code)
	}
	err = limberr.SetCustom(err, custom)
	glog.LogError(err, message, params...)
	return err
}

// TickValidate is automatically add validation error custom to the error
func TickValidate(err error, code string, message string, params ...interface{}) error {
	return TickCustom(err, ValidationFailedErr, code, message, params...)
}
