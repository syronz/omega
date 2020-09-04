package limberr

import (
	"fmt"
)

// WithMessage keeps the message of the error, each error can have one message
type WithMessage struct {
	Err error
	Msg string
}

func (p *WithMessage) Error() string { return fmt.Sprint(p.Err) }

// AddMessage split error in two parts
func AddMessage(err error, msg string) error {
	return &WithMessage{
		Err: err,
		Msg: msg,
	}
}

// WithCode is used for carrying the code of error
type WithCode struct {
	Err  error
	Code string
}

func (p *WithCode) Error() string { return fmt.Sprint(p.Err) }

//AddCode attach code to the error
func AddCode(err error, code string) error {
	return &WithCode{
		Err: fmt.Errorf("#%v, %w", code, err),
		// Err:  err,
		Code: code,
	}
}

type WithType struct {
	Err   error
	Type  string
	Title string
}

func (p *WithType) Error() string { return fmt.Sprint(p.Err) }

func AddType(err error, errType string, title string) error {
	return &WithType{
		Err:   err,
		Type:  errType,
		Title: title,
	}
}

// WithPath attach path to the error
type WithPath struct {
	Err  error
	Path string
}

func (p *WithPath) Error() string { return fmt.Sprint(p.Err) }

func AddPath(err error, path string) error {
	return &WithPath{
		Err:  err,
		Path: path,
	}
}

// WithStatus attach status to the error
type WithStatus struct {
	Err    error
	Status int
}

func (p *WithStatus) Error() string { return fmt.Sprint(p.Err) }

func AddStatus(err error, status int) error {
	return &WithStatus{
		Err:    err,
		Status: status,
	}
}

// WithDomain attach domain to the error
type WithDomain struct {
	Err    error
	Domain string
}

func (p *WithDomain) Error() string { return fmt.Sprint(p.Err) }

func AddDomain(err error, domain string) error {
	return &WithDomain{
		Err:    err,
		Domain: domain,
	}
}
