package engine

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"omega/config"
)

// Engine to keep all database connections and
// logs configuration and environments and etc
type Engine struct {
	DB           *gorm.DB
	ActivityDB   *gorm.DB
	ServerLog    *logrus.Logger
	ApiLog       *logrus.Logger
	Environments config.Environment
}

// Debug print struct with details with logrus ability
func (e *Engine) Debug(objs ...interface{}) {
	for _, v := range objs {
		e.ServerLog.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}

// DumpError print struct with details with logrus ability
func (e *Engine) DumpError(objs ...interface{}) {
	for _, v := range objs {
		e.ServerLog.Error(fmt.Sprintf("%T :: %+[1]v", v))
	}
}

// CheckError print all errors which happened inside the services, mainly they just have
// an error and a message
func (e *Engine) CheckError(err error, message string) {
	if err != nil {
		e.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(message)
	}
}

// CheckInfo print all errors which happened inside the services, mainly they just have
// an error and a message
func (e *Engine) CheckInfo(err error, message string) {
	if err != nil {
		e.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Info(message)
	}
}
