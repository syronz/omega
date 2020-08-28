package glog

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(format, output, level string, indent, file bool) {
	serverLogParam := LogParam{
		format:       format,
		output:       output,
		level:        level,
		JSONIndent:   indent,
		showFileLine: file, // true means filename and line number should be printed
	}

	logger = initLog(serverLogParam)
}

func New(format, output, level string, indent, file bool) *logrus.Logger {
	serverLogParam := LogParam{
		format:       format,
		output:       output,
		level:        level,
		JSONIndent:   indent,
		showFileLine: file, // true means filename and line number should be printed
	}

	return initLog(serverLogParam)
}

func initLog(p LogParam) *logrus.Logger {
	log := logrus.New()

	if p.showFileLine {
		hook := NewHook()
		hook.Field = "file"
		log.AddHook(hook)
	}

	setFormat(log, p)
	setOutput(log, p)
	setLevel(log, p)

	return log
}

func setFormat(log *logrus.Logger, p LogParam) {
	// TODO: should be completed
	switch p.format {
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{
			PrettyPrint: p.JSONIndent,
		})
	}
}

func setOutput(log *logrus.Logger, p LogParam) {
	switch p.output {
	case "stdout":
		log.SetOutput(os.Stdout)
	default:
		file, err := os.OpenFile(p.output, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err == nil {
			log.Out = file
		} else {
			log.Fatalf("Failed to write logs to file %q, [core/logs.go]", p.output)
		}
	}
}

func setLevel(log *logrus.Logger, p LogParam) {

	switch p.level {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "warn":
		log.SetLevel(logrus.WarnLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	}
}
