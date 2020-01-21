package server

import (
	"net/http"
	"omega/engine"
	"omega/internal/response"
	"omega/middleware"

	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine, e engine.Engine) {

	// static files
	r.Static("/public", "./public")

	// Root "/"
	routeRoot(r)

	// No Route "Not Found"
	routeNotFound(r)

	api := r.Group("/api/omega/v1")
	{
		routeAuth(api, e)

		api.Use(middleware.CheckToken(e))
		routeUser(api, e)
		routeActivity(api, e)
		routeRole(api, e)
	}

}

func routeRoot(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Result{
			Message: "Omega API Server v1.0",
		})
	})
}

func routeNotFound(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Result{
			Message: "Not Found",
		})
	})
}

func routeUser(api *gin.RouterGroup, e engine.Engine) {
	userAPI := initUserAPI(e)
	api.GET("/all/users", userAPI.FindAll)
	api.GET("/users", userAPI.List)
	api.GET("/users/:id", userAPI.FindByID)
	// api.POST("/users", userAPI.Create)
	api.POST("/users", userAPI.BuildCreate)
	api.PUT("/users/:id", userAPI.Update)
	api.DELETE("/users/:id", userAPI.Delete)
}

func routeAuth(api *gin.RouterGroup, e engine.Engine) {
	authAPI := initAuthAPI(e)
	api.POST("/auth/login", authAPI.Login)
}

func routeActivity(api *gin.RouterGroup, e engine.Engine) {
	activityAPI := initActivityAPI(e)
	api.GET("/activities", activityAPI.List)
	api.GET("/activities/:id", activityAPI.FindByID)
}

func routeRole(api *gin.RouterGroup, e engine.Engine) {
	roleAPI := initRoleAPI(e)
	api.GET("/all/roles", roleAPI.FindAll)
	api.GET("/roles", roleAPI.List)
	api.GET("/roles/:id", roleAPI.FindByID)
	api.POST("/roles", roleAPI.Create)
	api.PUT("/roles/:id", roleAPI.Update)
	api.DELETE("/roles/:id", roleAPI.Delete)
}
