package core

import (
	// Import Mysql and Postgress for testing
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"omega/engine"
	"omega/internal/core/setup"
	"omega/utils/glog"
)

// StartEngine initiate all different parts like logs and database connection and generate cfg
func StartEngine(models ...interface{}) (engine engine.Engine) {

	engine.Environments = getRegularEnvs()

	env := engine.Environments

	engine.ServerLog = setup.ServerLog(env)
	glog.Glog.ServerLog = engine.ServerLog

	engine.ApiLog = setup.APILog(env)
	glog.Glog.ApiLog = engine.ApiLog

	engine.DB = initDB(engine, env.Database.Data.Type, env.Database.Data.DSN, models...)
	engine.ActivityDB = initDB(engine, env.Database.Activity.Type, env.Database.Activity.DSN, models...)

	return
}
