package glog

import (
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Hook is a struct for hold main fields
type Hook struct {
	Field  string
	Skip   int
	levels []logrus.Level
}

// Levels return all levels of the hook
func (hook *Hook) Levels() []logrus.Level {
	return hook.levels
}

// Fire add new data to the entry.Data
func (hook *Hook) Fire(entry *logrus.Entry) error {
	// entry.Data[hook.Field] = hook.Formatter(findCaller(hook.Skip))
	entry.Data[hook.Field] = hook.formatter(findCaller(hook.Skip))
	var it interface{}

	err := json.Unmarshal([]byte(entry.Message), &it)
	if err == nil {
		entry.Data["data"] = it
		entry.Message = ""
	}

	return nil
}

func (hook *Hook) formatter(file, function string, line int) string {
	return fmt.Sprintf("%s:%d", file, line)
}
