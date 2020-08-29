package corerr

import (
	"fmt"
	"omega/internal/param"
	"omega/pkg/dict"
	"omega/pkg/glog"
)

type CustomError struct {
	Code          string      `json:"code,omitempty"`
	Type          string      `json:"type,omitempty"`
	Title         string      `json:"title,omitempty"`
	Domain        string      `json:"domain,omitempty"`
	Message       string      `json:"message,omitempty"`
	MessageParams []string    `json:"message_params,omitempty"`
	Path          string      `json:"path,omitempty"`
	InvalidParams interface{} `json:"invalid_params,omitempty"`
	Lang          dict.Lang   `json:"-"`
	Status        int         `json:"-"`
	ErrPanel      string      `json:"error_panel,omitempty"`
	OriginalError string      `json:"original_error,omitempty"`
}

func (p CustomError) Error() string {
	return fmt.Sprintf("custom error with code %v", p.Code)
}

// New is used for initiating an error
func New(code string, params param.Param, domain string, err error, data ...interface{}) CustomError {
	glog.Error(fmt.Sprintf("%v: %v, ", code, err), data)
	return CustomError{
		Code:     code,
		Domain:   domain,
		Lang:     params.Lang,
		ErrPanel: params.ErrPanel,
		// OriginalError:      fmt.Errorf("%w with data %v", err, data),
		OriginalError: err.Error(),
	}
}
