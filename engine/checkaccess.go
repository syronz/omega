package engine

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// CheckAccess compare the given resource with list of user's resources
func (e *Engine) CheckAccess(c *gin.Context, resource string) bool {
	var userID uint64

	if userIDtmp, ok := c.Get("USER_ID"); ok {
		userID = userIDtmp.(uint64)
	}

	resources := struct {
		Resources string
	}{}

	_ = e.DB.Table("users").Select("roles.resources").
		Joins("INNER JOIN roles ON users.role_id = roles.id").
		Where("users.id = ?", userID).Scan(&resources).Error

	result := strings.Contains(resources.Resources, resource)

	e.Debug("!!!!!!!!!!!!+++++++++++!!!!!!!!!!!!", result, resources.Resources, resource)
	return result

}
