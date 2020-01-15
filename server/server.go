package server

import (
	"omega/internal/core"
	"omega/middleware"

	"github.com/gin-gonic/gin"
)

// Setup integrate middleware and static route finally initiate router
func Setup(e core.Engine) *gin.Engine {

	r := gin.Default()

	e.LogAPI.Info("Server Started!")

	r.Use(middleware.APILogger())
	router(r, e)

	return r
}
