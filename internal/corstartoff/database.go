package corstartoff

import (
	"log"
	"omega/internal/core"

	"github.com/jinzhu/gorm"
)

// ConnectDB initiate the db connection by getting help from gorm
func ConnectDB(engine *core.Engine, printQueries bool) {
	var err error
	engine.DB, err = gorm.Open(engine.Envs[core.DatabaseDataType], engine.Envs[core.DatabaseDataDSN])
	if err != nil {
		log.Fatalln(err.Error())
	}

	engine.DB.LogMode(engine.Envs.ToBool(core.DatabaseDataLog))

	if printQueries {
		engine.DB.LogMode(true)
	}
}

// ConnectActivityDB initiate the db connection by getting help from gorm
func ConnectActivityDB(engine *core.Engine) {
	var err error
	engine.ActivityDB, err = gorm.Open(engine.Envs[core.DatabaseActivityType],
		engine.Envs[core.DatabaseActivityDSN])
	if err != nil {
		log.Fatalln(err.Error())
	}

	engine.ActivityDB.LogMode(engine.Envs.ToBool(core.DatabaseActivityLog))

}
