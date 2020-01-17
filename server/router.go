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
		res := response.Response{Context:c}
		res.Success(http.StatusOK, "Omega API Server v1.0", "", 0)
	})

	// No Route
	r.NoRoute(func(c *gin.Context) {
		res := response.Response{Context:c}
		res.Failed(http.StatusNotFound, 1404, "not found", "")
	})


	api := r.Group("/api/omega/v1")
	{
		userAPI := initUserAPI(e)
		api.GET("/users", userAPI.FindAll)
		api.GET("/user/:id", userAPI.FindByID)
		api.POST("/user", userAPI.Create)
		api.PUT("/user/:id", userAPI.Update)
		api.DELETE("/user/:id", userAPI.Delete)

		authAPI := initAuthAPI(e)
		api.POST("/auth/login", authAPI.Login)
	}

}
