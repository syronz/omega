package response

import (
	"omega/internal/core/corerr"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// NotBind use special custom_error for reduced it
func (r *Response) NotBind(err error, code, domain string, part string) {
	err = limberr.Take(err, code).Domain(domain).
		Message(corerr.ErrorInBindingV, dict.R(part)).
		Custom(corerr.BindingErr).Build()

	r.Error(err).JSON()
}

// Bind is used to make it more easear for binding items
func (r *Response) Bind(st interface{}, code, domain, part string) (err error) {
	if err = r.Context.ShouldBindJSON(&st); err != nil {
		r.NotBind(err, code, domain, part)
		return
	}

	return
}
