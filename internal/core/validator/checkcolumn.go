package validator

import (
	"omega/internal/core/corerr"
	"omega/pkg/helper"
	"omega/pkg/limberr"
	"strings"
)

// CheckColumns will check columns for security
func CheckColumns(cols []string, requestedCols string) (string, error) {
	var err error

	if requestedCols == "*" {
		return strings.Join(cols, ","), nil
	}

	variates := strings.Split(requestedCols, ",")
	for _, v := range variates {
		if ok, _ := helper.Includes(cols, v); !ok {
			err = limberr.AddInvalidParam(err, v, corerr.VisNotValid, v)
			err = limberr.SetCustom(err, corerr.ValidationFailedErr)
		}
	}

	return requestedCols, err

}
