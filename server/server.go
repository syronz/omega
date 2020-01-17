package server

import (
	"github.com/gin-gonic/gin"
	"omega/engine"
	"omega/middleware"
)

// Initialize integrate middleware and
// static route finally initiate router
func Initialize(e engine.Engine) *gin.Engine {
	r := gin.Default()
	r.Use(middleware.APILogger())
	router(r, e)
	return r
}
