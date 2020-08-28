package corerr

import (
	"fmt"
	"omega/pkg/dict"
	"omega/pkg/glog"
)

type CustomError struct {
	Code          string      `json:"code,omitempty"`
	Type          string      `json:"type,omitempty"`
	Title         string      `json:"title,omitempty"`
	Domain        string      `json:"domain,omitempty"`
	Message       string      `json:"message,omitempty"`
	Path          string      `json:"path,omitempty"`
	InvalidParams interface{} `json:"invalid_params,omitempty"`
}

func (p CustomError) Error() string {
	return fmt.Sprintf("custom error with code %v", p.Code)
}

// New is used for initiating an error
func New(lang dict.Language, domain, code string, err error, data interface{}) CustomError {
	glog.Error(err, data)
	return CustomError{
		Code:   code,
		Domain: domain,
	}
}
