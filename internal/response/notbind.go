package response

import (
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/internal/types"
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

// GetRowID convert string to the rowID and if not converted print a proper message
func (r *Response) GetRowID(idIn, code, part string) (id types.RowID, err error) {
	if id, err = types.StrToRowID(idIn); err != nil {
		err = limberr.Take(err, code).
			Message(corerr.InvalidVForV, dict.R(corterm.ID), dict.R(part)).
			Custom(corerr.ValidationFailedErr).Build()
		r.Error(err).JSON()
		return
	}

	return

}
