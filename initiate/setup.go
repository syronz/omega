package initiate

import (
	"github.com/kelseyhightower/envconfig"
	"log"
	"rest-gin-gorm/config"
)

// Setup initiate all difirent parts like log and database connection and generate cfg
func Setup() (cfg config.CFG) {
	var env config.Environment
	if err := envconfig.Process("server", &env); err != nil {
		log.Fatalln(err)
	}

	cfg.Log = initLog(env.Log.Format, env.Log.Output, env.Log.Level)

	if env.Database.Type == "" || env.Database.URL == "" {
		cfg.Log.Fatal("Environment is not set for database")
	}

	cfg.DB = initDB(env.Database.Type, env.Database.URL)

	return
}
