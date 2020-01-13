package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	// "rest-gin-gorm/tools/log"
)

func Setup(db *gorm.DB, log *logrus.Logger) *gin.Engine {
	log.Info("this is log >>>>>>>>>>>>>!!!!!!!!!!!!!!!!!!!")

	r := gin.Default()

	router(r, db, log)
	return r
}
