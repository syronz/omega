package config

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

type CFG struct {
	DB  *gorm.DB
	Log *logrus.Logger
}
