package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// CFG is used as a global struct but it is injected
type CFG struct {
	DB     *gorm.DB
	Log    *logrus.Logger
	Logapi *logrus.Logger
}

func (c *CFG) Debug(obj interface{}) {
	c.Log.Debug(fmt.Sprintf("%+v", obj))
}
