package corerr

import "strings"

// ClearDbErr find out what type of errors happened: duplicate, foreing keys or internal error
func ClearDbErr(err error) string {

	if strings.Contains(strings.ToUpper(err.Error()), "FOREIGN") {
		return "foreign"
	}
	if strings.Contains(strings.ToUpper(err.Error()), "DUPLICATE") {
		return "duplicate"
	}

	return "internal"

}
