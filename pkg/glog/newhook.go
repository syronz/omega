package glog

import (
	"github.com/sirupsen/logrus"
)

// NewHook is used for initiate an object of the Hook
func NewHook(levels ...logrus.Level) *Hook {
	hook := Hook{
		Field:  "source",
		Skip:   5,
		levels: levels,
	}
	if len(hook.levels) == 0 {
		hook.levels = logrus.AllLevels
	}

	return &hook
}
