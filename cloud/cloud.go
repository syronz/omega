package cloud

import (
	// "github.com/gin-contrib/cors"
	"net/http"
	"omega/engine"
	"omega/internal/response"
	"omega/middleware"

	"github.com/gin-gonic/gin"
)

// Initialize integrate middleware and
// static route finally initiate router
func Initialize(e engine.Engine) *gin.Engine {
	r := gin.Default()

	//r.Use(cors.New(cors.Config{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	//	AllowHeaders:     []string{"*"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//	AllowOriginFunc: func(origin string) bool {
	//		return origin == "http://erp14.com"
	//	},
	//	//MaxAge: 12 * time.Hour,
	//}))

	r.GET("/api/cloud/v1", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Result{
			Message: "Omega API Cloud v1.0",
		})
	})

	routeRoot(r)

	r.Use(middleware.APILogger())

	// api := r.Group("/api/cloud/v1")
	// {
	// }

	// router(r, e)
	return r
}

func routeRoot(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, response.Result{
			Message: "Omega API Cloud v1.0",
		})
	})
}
