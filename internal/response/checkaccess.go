package response

import (
	"omega/domain/base/basrepo"
	"omega/domain/service"
)

// CheckAccess is a helper for checking the permission for each method
func (r *Response) CheckAccess(resource string) bool {
	accessService := service.ProvideAccessService(basrepo.ProvideAccessRepo(r.Engine))

	return accessService.CheckAccess(r.Context, resource)
}
