package initiate

import (
	envEngine "github.com/caarlos0/env/v6"
	"log"
	"omega/internal/core"
	"omega/internal/glog"
)

// Setup initiate all different parts like log and database connection and generate cfg
func Setup() (engine core.Engine) {
	if err := envEngine.Parse(&engine.Environments); err != nil {
		log.Fatalln(err)
	}

	// cfg.Environments = env

	logParam := LogParam{
		format:       engine.Environments.Log.Format,
		output:       engine.Environments.Log.Output,
		level:        engine.Environments.Log.Level,
		showFileLine: true, // true means filename and line number should be printed
	}
	engine.Log = initLog(logParam)
	glog.GlobalLog.Logrus = engine.Log

	logAPIParam := LogParam{
		format:       engine.Environments.Logapi.Format,
		output:       engine.Environments.Logapi.Output,
		level:        engine.Environments.Logapi.Level,
		showFileLine: false,
	}
	engine.LogAPI = initLog(logAPIParam)
	glog.GlobalLog.Logapi = engine.LogAPI

	if engine.Environments.Database.Type == "" || engine.Environments.Database.URL == "" {
		engine.Log.Warn(engine.Environments)
		engine.Log.Fatal("Environment is not set for database")
	}

	engine.DB = initDB(engine, engine.Environments.Database.Type, engine.Environments.Database.URL)

	return
}
