package core

import (
	"log"
	"omega/config"
	"omega/engine"
	"omega/internal/glog"

	envEngine "github.com/caarlos0/env/v6"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
)

// StartEngine initiate all different parts like logs and database connection and generate cfg
func StartEngine() (engine engine.Engine) {
	if err := envEngine.Parse(&engine.Environments); err != nil {
		log.Fatalln(err)
	}

	env := engine.Environments

	engine.ServerLog = setupServerLog(env)
	glog.Glog.ServerLog = engine.ServerLog

	engine.ApiLog = setUpAPILog(env)
	glog.Glog.ApiLog = engine.ApiLog

	engine.DB = initDB(engine, env.Database.Data.Type, env.Database.Data.DSN)

	return
}

func setupServerLog(env config.Environment) *logrus.Logger {
	// Server logs's params
	serverLogParams := LogParam{
		format:       env.Log.ServerLog.Format,
		output:       env.Log.ServerLog.Output,
		level:        env.Log.ServerLog.Level,
		showFileLine: true, // true means filename and line number should be printed
	}

	return initLog(serverLogParams)
}

func setUpAPILog(env config.Environment) *logrus.Logger {
	// API logs's params
	apiLogParams := LogParam{
		format:       env.Log.ApiLog.Format,
		output:       env.Log.ApiLog.Output,
		level:        env.Log.ApiLog.Level,
		showFileLine: false,
	}

	return initLog(apiLogParams)
}
