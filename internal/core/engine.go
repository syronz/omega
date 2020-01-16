package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"omega/config"
)

// Engine is keep all database connections and log configuration and environments and etc
type Engine struct {
	DB           *gorm.DB
	ActivityDB   *gorm.DB
	Log          *logrus.Logger
	LogAPI       *logrus.Logger
	Environments config.Environment
}

// Debug print struct with details with logrus ability
func (e *Engine) Debug(objs ...interface{}) {
	for _, v := range objs {
		e.Log.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}

// Debug print struct with details with logrus ability
func (e *Engine) Error(objs ...interface{}) {
	for _, v := range objs {
		e.Log.Error(fmt.Sprintf("%T :: %+[1]v", v))
	}
}
