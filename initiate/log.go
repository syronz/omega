package initiate

import (
	"os"

	"github.com/sirupsen/logrus"
	"omega/internal/loghook"
)

func initLog(p LogParam) *logrus.Logger {

	// TODO: switch for each case should be completed

	log := logrus.New()

	if p.showFileLine {
		hook := loghook.NewHook()
		hook.Field = "file"
		log.AddHook(hook)
	}

	switch p.format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	switch p.output {
	case "stdout":
		log.SetOutput(os.Stdout)
	default:
		file, err := os.OpenFile(p.output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Fatalln("Failed to log to file", p.output)
		}
	}

	switch p.level {
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	}

	return log
}
