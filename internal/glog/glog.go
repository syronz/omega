// Package glog is stand for global log, it is not possible to inject log to all fuctions therefore
// we should import glog as a package
package glog

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// GLog is main struct for this package
type globalLog struct {
	Logrus *logrus.Logger
	Logapi *logrus.Logger
}

// GlobalLog is a global varable for initiate logrus
var GlobalLog globalLog

// Debug print struct with details with logrus ability
func Debug(objs ...interface{}) {
	for _, v := range objs {
		GlobalLog.Logrus.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}
