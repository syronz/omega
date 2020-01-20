package core

import (
	"log"
	"omega/engine"
	"omega/internal/core/setup"
	"omega/internal/glog"
	"omega/internal/models"
	"omega/pkg/role"
	"omega/pkg/user"

	envEngine "github.com/caarlos0/env/v6"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
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

	engine.DB = initDB(engine, env.Database.Data.Type, env.Database.Data.DSN,
		&user.User{},
		&role.Role{},
	)
	engine.ActivityDB = initDB(engine, env.Database.Activity.Type, env.Database.Activity.DSN,
		&models.Activity{})

	return
}
