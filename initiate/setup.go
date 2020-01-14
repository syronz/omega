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
	cfg.Log = initLog(env.Log.Format, env.Log.Output, env.Log.Level, true)
	cfg.Logapi = initLog(env.Logapi.Format, env.Logapi.Output, env.Logapi.Level, false)

	if env.Database.Type == "" || env.Database.URL == "" {
		cfg.Log.Warn(env)
		cfg.Log.Fatal("Environment is not set for database")
	}

	cfg.DB = initDB(env.Database.Type, env.Database.URL)

	return
}
