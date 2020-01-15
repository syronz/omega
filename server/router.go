package server

import (
	"github.com/gin-gonic/gin"
	"omega/internal/core"
)

func router(r *gin.Engine, e core.Engine) {

	userAPI := initUserAPI(e)
	r.GET("/users", userAPI.FindAll)
	r.GET("/users/:id", userAPI.FindByID)
	r.POST("/users", userAPI.Create)
	r.PUT("/users/:id", userAPI.Update)
	r.DELETE("/users/:id", userAPI.Delete)

}
