package limberr

import (
	"errors"
	"fmt"
)

// WithCode is used for carrying the code of error
type WithCode struct {
	Err  error
	Code string
}

func (p *WithCode) Error() string { return fmt.Sprint(p.Err) }

func AddCode(err error, code string) error {
	return &WithCode{
		Err:  fmt.Errorf("#%v, %w", code, err),
		Code: code,
	}
}

// WithMessage keeps the message of the error, each error can have one message
type WithMessage struct {
	Err    error
	Msg    string
	Params []interface{}
}

func (p *WithMessage) Error() string { return fmt.Sprint(p.Err) }

func AddMessage(err error, msg string, params ...interface{}) error {
	return &WithMessage{
		Err:    err,
		Msg:    msg,
		Params: params,
	}
}

// WithType is add type and title to the error
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

// WithInvalidParam holds invalid parameters
type WithInvalidParam struct {
	Err    error
	Field  string
	Reason string
	Params []interface{}
}

func (p *WithInvalidParam) Error() string { return fmt.Sprint(p.Err) }

func AddInvalidParam(err error, field, reason string, params ...interface{}) error {
	var gErr error
	if err == nil {
		gErr = errors.New(fmt.Sprintf(reason, params...))
	} else {
		gErr = err
	}

	return &WithInvalidParam{
		Err:    gErr,
		Field:  field,
		Reason: reason,
		Params: params,
	}
}

// WithCustom is used for holding the uniqError for filling the type and title based on local
// customization
type WithCustom struct {
	Err    error
	Custom CustomError
}

func (p *WithCustom) Error() string { return fmt.Sprint(p.Err) }

func SetCustom(err error, custom CustomError) error {
	return &WithCustom{
		Err:    err,
		Custom: custom,
	}
}
