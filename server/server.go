package server

import (
	"fmt"
	"omega/config"
	"omega/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Setup integrate middleware and static route finally initiate router
func Setup(c config.CFG) *gin.Engine {

	r := gin.Default()

	var tmp = struct {
		name string
		age  int
	}{
		"diako",
		23,
	}
	c.Log.Debug(tmp)
	c.Debug(tmp)
	c.Logapi.WithFields(logrus.Fields{
		"user":  fmt.Sprintf("%+v", tmp),
		"user2": tmp,
		// "name": "diako",
	}).Info("found user+++++++++++++++++++++++++")
	c.Log.Debug("this is here debug !!!!!!!!!!!!!!")
	c.Log.Info("this is here info !!!!!!!!!!!!!!")

	// r.Use(middleware.Logger(logrus.New()))

	// r.Use(middleware.Logger(c))
	// r.Use(middleware.GinBodyLogMiddleware)
	r.Use(middleware.Wrapper(c))
	router(r, c)

	return r
}
