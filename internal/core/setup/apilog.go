package setup

import (
	"omega/config"

	"github.com/sirupsen/logrus"
)

// APILog is used inside internal.core.StartEngine and test.core.StartEngine
func APILog(env config.Environment) *logrus.Logger {
	// API logs's params
	apiLogParams := LogParam{
		format:       env.Log.APILog.Format,
		output:       env.Log.APILog.Output,
		level:        env.Log.APILog.Level,
		JSONIndent:   env.Log.APILog.JSONIndent,
		showFileLine: false,
	}

	return initLog(apiLogParams)
}
