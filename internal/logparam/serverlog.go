package logparam

import (
	"omega/internal/core"
)

// ServerLog connected to the engine
func ServerLog(engine *core.Engine) {
	// Server logs's params
	serverLogParam := LogParam{
		format:       engine.Envs[core.ServerLogFormat],
		output:       engine.Envs[core.ServerLogOutput],
		level:        engine.Envs[core.ServerLogLevel],
		JSONIndent:   engine.Envs.ToBool(core.ServerLogJSONIndent),
		showFileLine: true, // true means filename and line number should be printed
	}

	engine.ServerLog = initLog(serverLogParam)

}
