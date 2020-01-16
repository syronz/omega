package server

import (
	"github.com/gin-gonic/gin"
	"omega/internal/core"
)

func router(r *gin.Engine, e core.Engine) {
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
