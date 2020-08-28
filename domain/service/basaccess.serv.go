package service

import (
	"omega/domain/base/basrepo"
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

var thisCache map[types.RowID]string

func init() {
	thisCache = make(map[types.RowID]string)
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

	if resources, ok = thisCache[userID]; !ok {
		var err error
		resources, err = p.Repo.GetUserResources(userID)
		glog.CheckError(err, "error in finding the resources for user", userID)
		BasAccessAddToCache(userID, resources)
	}

	return !strings.Contains(resources, string(resource))

}

// BasAccessAddToCache add the resources to the thisCache
func BasAccessAddToCache(userID types.RowID, resources string) {
	thisCache[userID] = resources
}

func BasAccessDeleteFromCache(userID types.RowID) {
	delete(thisCache, userID)
}

func BasAccessResetCache(userID types.RowID) {
	thisCache[userID] = ""
}

func BasAccessResetFullCache() {
	thisCache = make(map[types.RowID]string)
}
