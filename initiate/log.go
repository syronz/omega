package initiate

import (
	"os"

	"github.com/sirupsen/logrus"
)

func initLog(format, output, level string) *logrus.Logger {
	// TODO: switch for each case should be completed

	log := logrus.New()
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
	}

	return log
}
