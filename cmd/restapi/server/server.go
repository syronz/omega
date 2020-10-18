package server

import (
	"fmt"
	"log"
	"net/http"
	"omega/domain/base"
	"omega/internal/core"
	"omega/internal/core/corerr"
	"omega/internal/core/cormid"
	"omega/internal/response"
	"omega/pkg/glog"
	"omega/pkg/limberr"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// Start initiate the server
func Start(engine *core.Engine) *gin.Engine {

	var r *gin.Engine
	if engine.Envs[core.GindMode] == "debug" {
		r = gin.Default()
	} else {
		r = gin.New()
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "127.0.0.1"
		},
		//MaxAge: 12 * time.Hour,
	}))
	r.Use(cormid.APILogger(engine))

	// No Route "Not Found"
	notFoundRoute(r, engine)

	rg := r.Group("/api/restapi/v1")
	{
		Route(*rg, engine)
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%v:%v", engine.Envs[core.Addr], engine.Envs[core.Port]),
		Handler: r,
		//TLSEnvironment:    tlsEnvironment,
		//TLSEnvironment:    nil,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  360 * time.Second,
	}

	glog.Info("Rest-API starting server on ", engine.Envs[core.Addr], ":", engine.Envs[core.Port], "***********************************************************************")
	fmt.Printf("Rest-API starting server on %v:%v\n", engine.Envs[core.Addr], engine.Envs[core.Port])
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}

	return r
}

func notFoundRoute(r *gin.Engine, engine *core.Engine) {
	r.NoRoute(func(c *gin.Context) {
		err := limberr.New("route not found", "E1015777").Custom(corerr.RouteNotFoundErr).
			Message(corerr.PleaseReportErrorToProgrammer).Build()
		response.New(engine, c, base.Domain).Error(err).JSON()
	})
}
