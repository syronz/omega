package server

import (
	"rest-gin-gorm/config"

	"github.com/gin-gonic/gin"
)

// Setup integrate middleware and static route finally initiate router
func Setup(c config.CFG) *gin.Engine {

	r := gin.Default()

	router(r, c)
	return r
}
