package initiate

import (
	envEngine "github.com/caarlos0/env/v6"
	"log"
	"omega/config"
)

// Setup initiate all difirent parts like log and database connection and generate cfg
func Setup() (cfg config.CFG, env config.Environment) {
	if err := envEngine.Parse(&env); err != nil {
		log.Fatalln(err)
	}

	// the last true means filename and line number should be printed
	logParam := LogParam{
		format:  env.Log.Format,
		output:  env.Log.Output,
		level:   env.Log.Level,
		hasHook: true,
	}
	cfg.Log = initLog(logParam)

	logAPIParam := LogParam{
		format:  env.Logapi.Format,
		output:  env.Logapi.Output,
		level:   env.Logapi.Level,
		hasHook: false,
	}
	cfg.Logapi = initLog(logAPIParam)

	if env.Database.Type == "" || env.Database.URL == "" {
		cfg.Log.Warn(env)
		cfg.Log.Fatal("Environment is not set for database")
	}

	cfg.DB = initDB(cfg, env.Database.Type, env.Database.URL)

	return
}
