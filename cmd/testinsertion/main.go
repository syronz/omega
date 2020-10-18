package main

import (
	"flag"
	"omega/cmd/restapi/startoff"
	"omega/cmd/testinsertion/insertdata"
	"omega/internal/core"
	"omega/internal/corstartoff"
	"omega/pkg/dict"
	"omega/pkg/glog"
	"omega/test/kernel"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var noReset bool
var logQuery bool

func init() {
	flag.BoolVar(&noReset, "noReset", false, "by default it drop tables before migrate")
	flag.BoolVar(&logQuery, "logQuery", false, "print queries in gorm")
}

func main() {
	flag.Parse()

	engine := kernel.LoadTestEnv()

	glog.Init(engine.Envs[core.ServerLogFormat],
		engine.Envs[core.ServerLogOutput],
		engine.Envs[core.ServerLogLevel],
		engine.Envs.ToBool(core.ServerLogJSONIndent),
		true)

	dict.Init(engine.Envs[core.TermsPath], engine.Envs.ToBool(core.TranslateInBackend))

	corstartoff.ConnectDB(engine, logQuery)
	corstartoff.ConnectActivityDB(engine)
	startoff.Migrate(engine)
	insertdata.Insert(engine)

	if noReset {
		glog.Debug("Data has been migrated successfully (no reset)")
		// fmt.Println("Data has been migrated successfully (no reset)")
	} else {
		glog.Debug("Data has been reset successfully")
		// fmt.Println("Data has been reset successfully")
	}

}
