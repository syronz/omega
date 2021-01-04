package service

import (
	"omega/domain/base/basrepo"
	"omega/domain/sync"
	"omega/internal/core"
	"omega/internal/types"
	"omega/pkg/glog"
	"strings"

	"github.com/gin-gonic/gin"
)

// BasAccessServ defining auth service
type BasAccessServ struct {
	Repo   basrepo.AccessRepo
	Engine *core.Engine
}

// ProvideBasAccessService for auth is used in wire
func ProvideBasAccessService(p basrepo.AccessRepo) BasAccessServ {
	return BasAccessServ{Repo: p, Engine: p.Engine}
}

var cacheResource map[types.RowID]string

func init() {
	cacheResource = make(map[types.RowID]string)
}

// CheckAccess is used inside each method to findout if user has permission or not
func (p *BasAccessServ) CheckAccess(c *gin.Context, resource types.Resource) bool {
	var userID types.RowID

	if userIDtmp, ok := c.Get("USER_ID"); ok {
		userID = userIDtmp.(types.RowID)
	} else {
		return true
	}

	var resources string
	var ok bool

	if resources, ok = cacheResource[userID]; !ok {
		var err error
		resources, err = p.Repo.GetUserResources(userID)
		glog.CheckError(err, "error in finding the resources for user", userID)
		BasAccessAddToCache(userID, resources)
	}

	return !strings.Contains(resources, string(resource))

}

func IsSuperAdmin(userID types.RowID) bool {
	return strings.Contains(cacheResource[userID], string(sync.SuperAdmin))
}

// CheckRange is used for checking if user has access to special range of data
func (p *BasAccessServ) CheckRange(companyID, nodeID uint64) bool {
	if companyID > 0 {
		if companyID != 1001 {
			return false
		}
	}

	return true
}

// BasAccessAddToCache add the resources to the cacheResource
func BasAccessAddToCache(userID types.RowID, resources string) {
	cacheResource[userID] = resources
}

func BasAccessDeleteFromCache(userID types.RowID) {
	delete(cacheResource, userID)
}

func BasAccessResetCache(userID types.RowID) {
	cacheResource[userID] = ""
}

func BasAccessResetFullCache() {
	cacheResource = make(map[types.RowID]string)
}
