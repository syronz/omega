package core

import (
	"omega/engine"
	"omega/internal/core/setup"
	"omega/utils/glog"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	return
}
