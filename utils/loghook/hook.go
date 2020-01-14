package loghook

import (
	"github.com/sirupsen/logrus"
)

// Hook is a struct for hold main fields
type Hook struct {
	Field     string
	Skip      int
	levels    []logrus.Level
	Formatter func(file, function string, line int) string
}

// Levels return all levels of the hook
func (hook *Hook) Levels() []logrus.Level {
	return hook.levels
}

// Fire add new data to the entry.Data
func (hook *Hook) Fire(entry *logrus.Entry) error {
	entry.Data[hook.Field] = hook.Formatter(findCaller(hook.Skip))
	return nil
}
