package response

import (
	"omega/internal/core/corerr"
	"omega/pkg/limberr"
)

// GetCompanyNode return back node_id and company_id
func (r *Response) GetCompanyNode(code, domain string) (companyID, nodeID uint64, err error) {
	var ok bool
	var intCompanyID, intNodeID interface{}

	if intCompanyID, ok = r.Context.Get("COMPANY_ID"); !ok {
		err = limberr.New("company_id not exist in the token", code).
			Domain(domain).
			Message(corerr.VNotExist, "company_id").
			Custom(corerr.ForbiddenErr).Build()
		return
	}
	companyID = intCompanyID.(uint64)

	if intNodeID, ok = r.Context.Get("NODE_ID"); !ok {
		err = limberr.New("node_id not exist in the token", code).
			Domain(domain).
			Message("node_id not exist").
			Custom(corerr.ForbiddenErr).Build()
		return
	}
	nodeID = intNodeID.(uint64)

	// err = limberr.Take(err, code).Domain(domain).
	// 	Message(corerr.ErrorInBindingV, dict.R(part)).
	// 	Custom(corerr.Forbidden).Build()

	return
}
