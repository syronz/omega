package glog

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

// LogParam used for parameter between start and initLog
type LogParam struct {
	format       string
	output       string
	level        string
	JSONIndent   bool
	showFileLine bool
}
