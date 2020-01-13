package initiate

import (
	"fmt"
	"rest-gin-gorm/config"

	// "github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	"log"
)

func Setup() (cfg config.CFG) {
	var env config.Environment
	if err := envconfig.Process("server", &env); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\n\n+++++++++++++++ %+v \n\n", env)

	cfg.Log = initLog(env.Log.Format, env.Log.Output, env.Log.Level)

	if env.Database.Type == "" || env.Database.URL == "" {
		cfg.Log.Fatal("Environment is not set for database")
	}

	cfg.DB = initDB(env.Database.Type, env.Database.URL)

	return
}
