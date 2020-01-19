package setup

import (
	"omega/config"

	"github.com/sirupsen/logrus"
)

// ServerLog is used inside internal.core.StartEngine and test.core.StartEngine
func ServerLog(env config.Environment) *logrus.Logger {
	// Server logs's params
	serverLogParams := LogParam{
		format:       env.Log.ServerLog.Format,
		output:       env.Log.ServerLog.Output,
		level:        env.Log.ServerLog.Level,
		JSONIndent:   env.Log.ServerLog.JSONIndent,
		showFileLine: true, // true means filename and line number should be printed
	}

	return initLog(serverLogParams)
}
