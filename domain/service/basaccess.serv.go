package service

import (
	"omega/domain/base/basrepo"
	"omega/internal/core"
	"omega/internal/types"
	"strings"

	"github.com/gin-gonic/gin"
)

// AccessServ defining auth service
type AccessServ struct {
	Repo   basrepo.AccessRepo
	Engine *core.Engine
}

// ProvideAccessService for auth is used in wire
func ProvideAccessService(p basrepo.AccessRepo) AccessServ {
	return AccessServ{Repo: p, Engine: p.Engine}
}

var thisCache map[types.RowID]string

func init() {
	thisCache = make(map[types.RowID]string)
}

// CheckAccess is used inside each method to findout if user has permission or not
func (p *AccessServ) CheckAccess(c *gin.Context, resource string) bool {
	var userID types.RowID

	p.Engine.Debug(thisCache)
	// thisCache++

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
		p.Engine.CheckError(err, "error in finding the resources for user", userID)
		BasAccessAddToCache(userID, resources)
	}

	return !strings.Contains(resources, resource)

}

// AddResourceToCache add the resources to the thisCache
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
