package core

import "fmt"

// ErrorWithStatus is used in case of record not found
type ErrorWithStatus struct {
	Err     string `json:"error,omitempty"`
	Status  int
	Message string `json:"message,omitempty"`
}

func (p *ErrorWithStatus) Error() string {
	return fmt.Sprintf("%v", p.Err)
}

// NewErrorWithStatus initiate an error with status code
func NewErrorWithStatus(err string, status int) *ErrorWithStatus {
	notFoundError := ErrorWithStatus{Err: err, Status: status}
	return &notFoundError
}

// SetMsg set custom text instead of showing error, it automatically translate it
func (p *ErrorWithStatus) SetMsg(msg string) *ErrorWithStatus {
	p.Message = msg
	return p
}
