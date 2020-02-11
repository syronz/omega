package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"omega/engine"
	"omega/middleware"
)

// Initialize integrate middleware and
// static route finally initiate router
func Initialize(e engine.Engine) *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://erp14.com"
		},
		//MaxAge: 12 * time.Hour,
	}))

	r.Use(middleware.APILogger())
	router(r, e)
	return r
}
