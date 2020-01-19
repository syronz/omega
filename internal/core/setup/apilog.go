package setup

import (
	"omega/config"

	"github.com/sirupsen/logrus"
)

// APILog is used inside internal.core.StartEngine and test.core.StartEngine
func APILog(env config.Environment) *logrus.Logger {
	// API logs's params
	apiLogParams := LogParam{
		format:       env.Log.ApiLog.Format,
		output:       env.Log.ApiLog.Output,
		level:        env.Log.ApiLog.Level,
		JSONIndent:   env.Log.ApiLog.JSONIndent,
		showFileLine: false,
	}

	return initLog(apiLogParams)
}
