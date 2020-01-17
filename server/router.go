package server

import (
	"net/http"
	"omega/engine"
	"omega/internal/response"
	"omega/middleware"

	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine, e engine.Engine) {
	// Root "/"
	routeRoot(r)

	// No Route "Not Found"
	routeNotFound(r)

	api := r.Group("/api/omega/v1")
	{
		routeAuth(api, e)

		api.Use(middleware.CheckToken(e))
		routeUser(api, e)
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
	api.GET("/users", userAPI.FindAll)
	api.GET("/users/:id", userAPI.FindByID)
	api.POST("/users", userAPI.Create)
	api.PUT("/users/:id", userAPI.Update)
	api.DELETE("/users/:id", userAPI.Delete)
}

func routeAuth(api *gin.RouterGroup, e engine.Engine) {
	authAPI := initAuthAPI(e)
	api.POST("/auth/login", authAPI.Login)
}
