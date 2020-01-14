package server

import (
	"omega/config"
	"omega/middleware"

	"github.com/gin-gonic/gin"
)

// Setup integrate middleware and static route finally initiate router
func Setup(c config.CFG) *gin.Engine {

	r := gin.Default()

	c.Logapi.Info("Server Started!")

	r.Use(middleware.APILogger(c))
	router(r, c)

	return r
}
