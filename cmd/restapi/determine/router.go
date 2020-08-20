package determine

import (
	"omega/internal/core"
	"omega/internal/middleware"

	"github.com/gin-gonic/gin"
)

// Route trigger router and api methods
func Route(rg gin.RouterGroup, engine *core.Engine) {
	// Base Domain
	basAuthAPI := initBasAuthAPI(engine)
	basUserAPI := initBasUserAPI(engine)
	basRoleAPI := initBasRoleAPI(engine)
	basSettingAPI := initBasSettingAPI(engine)
	basActivityAPI := initBasActivityAPI(engine)

	rg.POST("/login", basAuthAPI.Login)

	rg.Use(middleware.AuthGuard(engine))

	rg.POST("/logout", basAuthAPI.Logout)

	// Base Domain
	rg.GET("/temporary/token", basAuthAPI.TemporaryToken)

	rg.GET("/settings", basSettingAPI.List)
	rg.GET("/settings/:settingID", basSettingAPI.FindByID)
	rg.PUT("/settings/:settingID", basSettingAPI.Update)

	rg.GET("/roles", basRoleAPI.List)
	rg.GET("/roles/:roleID", basRoleAPI.FindByID)
	rg.POST("/roles", basRoleAPI.Create)
	rg.PUT("/roles/:roleID", basRoleAPI.Update)
	rg.DELETE("/roles/:roleID", basRoleAPI.Delete)
	rg.GET("excel/roles", basRoleAPI.Excel)

	rg.GET("/username/:username", basUserAPI.FindByUsername)
	rg.GET("/users", basUserAPI.List)
	rg.GET("/users/:userID", basUserAPI.FindByID)
	rg.POST("/users", basUserAPI.Create)
	rg.PUT("/users/:userID", basUserAPI.Update)
	rg.DELETE("/users/:userID", basUserAPI.Delete)
	rg.GET("excel/users", basUserAPI.Excel)

	rg.GET("/activities", basActivityAPI.List)

}
