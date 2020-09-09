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
