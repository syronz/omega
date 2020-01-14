package initiate

import (
	"os"

	"github.com/sirupsen/logrus"
	"omega/utils/loghook"
)

func initLog(format, output, level string, hasHook bool) *logrus.Logger {
	// TODO: switch for each case should be completed

	log := logrus.New()

	if hasHook {
		hook := loghook.NewHook()
		hook.Field = "file"
		log.AddHook(hook)
	}

	switch format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	switch output {
	case "output":
		log.SetOutput(os.Stdout)
	}

	switch level {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	}

	return log
}
