package response

import (
	"omega/domain/base/basrepo"
	"omega/domain/service"
	"omega/internal/types"
)

// CheckAccess is a helper for checking the permission for each method
func (r *Response) CheckAccess(resource types.Resource) bool {
	accessService := service.ProvideAccessService(basrepo.ProvideAccessRepo(r.Engine))

	return accessService.CheckAccess(r.Context, resource)
}
