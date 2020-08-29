package response

import "omega/internal/core/corerr"

// NotBind use special custom_error for reduced it
func (r *Response) NotBind(code, domain string, err error, field, path string) {
	err = corerr.New(code, r.params, domain, err).
		NotBind(field, path)
	r.Error(err).JSON()
}
