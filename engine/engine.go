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
