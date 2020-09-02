package corerr

import (
	"fmt"
	"omega/internal/param"
	"omega/pkg/dict"
	"omega/pkg/glog"
)

type CustomError struct {
	Code          string        `json:"code,omitempty"`
	Type          string        `json:"type,omitempty"`
	Title         string        `json:"title,omitempty"`
	Domain        string        `json:"domain,omitempty"`
	Message       string        `json:"message,omitempty"`
	MessageParams []interface{} `json:"message_params,omitempty"`
	Path          string        `json:"path,omitempty"`
	InvalidParams []Field       `json:"invalid_params,omitempty"`
	Lang          dict.Lang     `json:"-"`
	Status        int           `json:"-"`
	ErrPanel      string        `json:"-"`
	OriginalError string        `json:"original_error,omitempty"`
	Data          []interface{} `json:"data,omitempty"`
}

// Field is used as an array inside the FieldError
type Field struct {
	Field        string        `json:"field,omitempty"`
	Reason       string        `json:"reason,omitempty"`
	ReasonParams []interface{} `json:"reason_params,omitempty"`
}

func (p CustomError) Error() string {
	errStr := fmt.Sprintf("custom error with code:%v, msg:%v", p.Code, p.Message)
	if len(p.InvalidParams) > 0 {
		errStr += fmt.Sprintf(", invalid_params:%+v", p.InvalidParams)
	}
	return errStr
}

// NewSilent is used for initiating an error
func NewSilent(code string, params param.Param, domain string, err error, data ...interface{}) CustomError {
	var errStr string
	if err != nil {
		errStr = err.Error()
	}

	return CustomError{
		Code:          code,
		Domain:        domain,
		Lang:          params.Lang,
		ErrPanel:      params.ErrPanel,
		OriginalError: errStr,
		Data:          data,
	}
}

// New is used for initiating an error
func New(code string, params param.Param, domain string, err error, data ...interface{}) CustomError {
	glog.Error(fmt.Sprintf("%v: %v, ", code, err), data)

	return NewSilent(code, params, domain, err, data...)
}