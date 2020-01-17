package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"omega/engine"
	"omega/internal/response"
)

func router(r *gin.Engine, e engine.Engine) {
	// Root
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Result{
			Status:  true,
			Message: "Omega API Server v1.0",
		})
	})

	// No Route
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, response.Result{
			Status:  false,
			Message: "Not Found",
			Code:    1404,
		})
	})

	api := r.Group("/api/omega/v1")
	{
		userAPI := initUserAPI(e)
		api.GET("/users", userAPI.FindAll)
		api.GET("/users/:id", userAPI.FindByID)
		api.POST("/users", userAPI.Create)
		api.PUT("/users/:id", userAPI.Update)
		api.DELETE("/users/:id", userAPI.Delete)

		authAPI := initAuthAPI(e)
		api.POST("/auth/login", authAPI.Login)
	}

}
