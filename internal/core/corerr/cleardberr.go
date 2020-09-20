package corerr

import (
	"omega/pkg/limberr"
	"strings"

	"github.com/jinzhu/gorm"
)

// ClearDbErr find out what type of errors happened: duplicate, foreign keys or internal error
func ClearDbErr(err error) limberr.CustomError {
	if err == nil {
		return Nil
	}

	if strings.Contains(strings.ToUpper(err.Error()), "FOREIGN") {
		return ForeignErr
	}
	if strings.Contains(strings.ToUpper(err.Error()), "DUPLICATE") {
		return DuplicateErr
	}
	if strings.Contains(strings.ToUpper(err.Error()), "UNKNOWN COLUMN") {
		return ValidationFailedErr
	}

	if gorm.IsRecordNotFoundError(err) {
		return NotFoundErr
	}

	return UnkownErr

}
