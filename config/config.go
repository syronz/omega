package config

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// CFG is used as a global struct but it is injected
type CFG struct {
	DB  *gorm.DB
	Log *logrus.Logger
}
