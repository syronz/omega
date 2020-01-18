// Package glog is stand for global logs, it is not possible to inject logs to all fuctions therefore
// we should import glog as a package
package glog

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

// Log is main struct for this package
type Log struct {
	ServerLog *logrus.Logger
	ApiLog    *logrus.Logger
}

// Glog is a global variable for initiate logrus
var Glog Log

// Debug print struct with details with logrus ability
func Debug(objs ...interface{}) {
	for _, v := range objs {
		Glog.ServerLog.Debug(fmt.Sprintf("%T :: %+[1]v", v))
	}
}

// CheckError print all errors which happened inside the services, mainly they just have
// an error and a message
func CheckError(err error, message string) {
	if err != nil {
		Glog.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Error(message)
	}
}

// CheckInfo print all errors which happened inside the services, mainly they just have
// an error and a message
func CheckInfo(err error, message string) {
	if err != nil {
		Glog.ServerLog.WithFields(logrus.Fields{
			"err": err.Error(),
		}).Info(message)
	}
}
