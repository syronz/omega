package core

import (
	envEngine "github.com/caarlos0/env/v6"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"omega/engine"
	"omega/internal/glog"
)

// StartEngine initiate all different parts like logs and database connection and generate cfg
func StartEngine() (engine engine.Engine) {
	if err := envEngine.Parse(&engine.Environments); err != nil {
		log.Fatalln(err)
	}

	env := engine.Environments

	// Server logs's params
	serverLogParams := LogParam {
		format:       env.ServerLog.Format,
		output:       env.ServerLog.Output,
		level:        env.ServerLog.Level,
		showFileLine: true, // true means filename and line number should be printed
	}
	engine.ServerLog = initLog(serverLogParams)
	glog.Glog.ServerLog = engine.ServerLog

	// API logs's params
	apiLogParams := LogParam {
		format:       env.ApiLog.Format,
		output:       env.ApiLog.Output,
		level:        env.ApiLog.Level,
		showFileLine: false,
	}
	engine.ApiLog = initLog(apiLogParams)
	glog.Glog.ApiLog = engine.ApiLog

	engine.DB = initDB(engine, env.Database.Type, env.Database.DSN)

	return
}
