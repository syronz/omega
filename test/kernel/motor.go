package kernel

import (
	"omega/internal/core"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// StartMotor for generating engine special for TDD
func StartMotor(printQueries bool, debugLevel bool) *core.Engine {
	engine := LoadTestEnv()
	// logparam.ServerLog(engine)
	// corstartoff.LoadTerms(engine)
	// if debugLevel {
	// 	engine.Envs[core.ServerLogLevel] = "trace"
	// }
	// logparam.ServerLog(engine)
	// corstartoff.ConnectDB(engine, printQueries)
	// corstartoff.ConnectActivityDB(engine)

	return engine
}
