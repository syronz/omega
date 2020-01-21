package core

import (
	"log"
	"omega/engine"
	"omega/internal/core/migrate"
	"omega/internal/core/setup"
	"omega/internal/glog"

	envEngine "github.com/caarlos0/env/v6"
)

// StartEngine initiate all different parts like logs and database connection and generate cfg
func StartEngine() (engine engine.Engine) {
	if err := envEngine.Parse(&engine.Environments); err != nil {
		log.Fatalln(err)
	}

	env := engine.Environments

	engine.ServerLog = setup.ServerLog(env)
	glog.Glog.ServerLog = engine.ServerLog

	engine.ApiLog = setup.APILog(env)
	glog.Glog.ApiLog = engine.ApiLog

	engine.DB = initDataDB(engine, env.Database.Data.Type, env.Database.Data.DSN)
	engine.ActivityDB = initActivityDB(engine, env.Database.Activity.Type, env.Database.Activity.DSN)

	migrate.InsertData(engine)

	return
}
