// Package glog is stand for global logs, it is not possible to inject logs to all fuctions therefore
// we should import glog as a package
package glog

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// GLog is main struct for this package
type Log struct {
	ServerLog *logrus.Logger
	ApiLog *logrus.Logger
}

// GlobalLog is a global variable for initiate logrus
var Glog Log

// Debug print struct with details with logrus ability
func Debug(objs ...interface{}) {
	for _, v := range objs {
		Glog.ServerLog.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}
