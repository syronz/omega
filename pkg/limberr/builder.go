package limberr

import "errors"

type LimbErr struct {
	err error
}

func New(errStr string, code ...string) *LimbErr {
	var limbErr LimbErr
	err := errors.New(errStr)

	if len(code) > 0 {
		limbErr.err = AddCode(err, code[0])
	} else {
		limbErr.err = err
	}

	return &limbErr

}

// Take initiate the
func Take(err error, code ...string) *LimbErr {
	var limbErr LimbErr
	if len(code) > 0 {
		limbErr.err = AddCode(err, code[0])
	} else {
		limbErr.err = err
	}

	return &limbErr
}

func (p *LimbErr) Code(code string) *LimbErr {
	p.err = AddCode(p.err, code)
	return p
}

func (p *LimbErr) Message(message string, params ...interface{}) *LimbErr {
	p.err = AddMessage(p.err, message, params...)
	return p
}

func (p *LimbErr) Custom(custom CustomError) *LimbErr {
	p.err = SetCustom(p.err, custom)
	return p
}

func (p *LimbErr) Domain(domain string) *LimbErr {
	p.err = AddDomain(p.err, domain)
	return p
}

func (p *LimbErr) Path(path string) *LimbErr {
	p.err = AddPath(p.err, path)
	return p
}

func (p *LimbErr) Status(status int) *LimbErr {
	p.err = AddStatus(p.err, status)
	return p
}

func (p *LimbErr) InvalidParam(field, reason string, params ...interface{}) *LimbErr {
	p.err = AddInvalidParam(p.err, field, reason, params...)
	return p
}

func (p *LimbErr) Build() error {
	return p.err
}
