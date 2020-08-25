package determine

import (
	"omega/domain/base"
	"omega/domain/base/basmid"
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
	access := basmid.NewAccessMid(engine)

	rg.POST("/logout", basAuthAPI.Logout)

	// Base Domain
	rg.GET("/temporary/token", basAuthAPI.TemporaryToken)

	rg.GET("/settings", access.Check(base.BasSettingRead), basSettingAPI.List)
	rg.GET("/settings/:settingID", access.Check(base.BasSettingRead), basSettingAPI.FindByID)
	rg.PUT("/settings/:settingID", access.Check(base.BasSettingWrite), basSettingAPI.Update)

	rg.GET("/roles", access.Check(base.BasRoleRead), basRoleAPI.List)
	rg.GET("/roles/:roleID", access.Check(base.BasRoleRead), basRoleAPI.FindByID)
	rg.POST("/roles", access.Check(base.BasRoleWrite), basRoleAPI.Create)
	rg.PUT("/roles/:roleID", access.Check(base.BasRoleWrite), basRoleAPI.Update)
	rg.DELETE("/roles/:roleID", access.Check(base.BasRoleWrite), basRoleAPI.Delete)
	rg.GET("excel/roles", access.Check(base.BasRoleExcel), basRoleAPI.Excel)

	rg.GET("/username/:username", access.Check(base.BasUserRead), basUserAPI.FindByUsername)
	rg.GET("/users", access.Check(base.BasUserRead), basUserAPI.List)
	rg.GET("/users/:userID", access.Check(base.BasUserRead), basUserAPI.FindByID)
	rg.POST("/users", access.Check(base.BasUserWrite), basUserAPI.Create)
	rg.PUT("/users/:userID", access.Check(base.BasUserWrite), basUserAPI.Update)
	rg.DELETE("/users/:userID", access.Check(base.BasUserWrite), basUserAPI.Delete)
	rg.GET("excel/users", access.Check(base.BasUserExcel), basUserAPI.Excel)

	rg.GET("/activities", access.Check(base.BasActivityAll), basActivityAPI.List)

	//TODO: delete below
	rg.GET("/ping", access.Check(base.BasPing), pong)

}

func pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "hello diako",
	})
}
