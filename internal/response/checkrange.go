package response

import (
	"fmt"
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/core/corerr"
	"omega/internal/core/corterm"
	"omega/pkg/dict"
	"omega/pkg/limberr"
)

// CheckRange will checks the range for companyID and nodeID
func (r *Response) CheckRange(companyID, nodeID uint64) bool {
	accessServ := service.ProvideBasAccessService(basrepo.ProvideAccessRepo(r.Engine))

	result := accessServ.CheckRange(companyID, nodeID)

	if !result {
		err := limberr.New("you don't have permission to this scope", "E1052722").
			Message(corerr.ForbiddenToVV, dict.R(corterm.Scope), fmt.Sprintf("%v %v", companyID, nodeID)).
			Custom(corerr.ForbiddenErr).Build()
		r.Error(err).JSON()
	}

	return result
}
