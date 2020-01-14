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
	case "stdout":
		log.SetOutput(os.Stdout)
	default:
		// log.SetOutput(os.Stdout)
		file, err := os.OpenFile(output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Fatalln("Failed to log to file", output)
		}
	}

	switch level {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	}

	return log
}
