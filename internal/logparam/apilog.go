package logparam

import (
	"omega/internal/core"
)

// APILog is used inside internal.core.StartEngine and test.core.StartEngine
func APILog(engine *core.Engine) {
	apiLogParams := LogParam{
		format:       engine.Envs[core.APILogFormat],
		output:       engine.Envs[core.APILogOutput],
		level:        engine.Envs[core.APILogLevel],
		JSONIndent:   engine.Envs.ToBool(core.APILogJSONIndent),
		showFileLine: false,
	}

	engine.APILog = initLog(apiLogParams)
}
