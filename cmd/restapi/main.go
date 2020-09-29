package main

import (
	"omega/cmd/restapi/insertdata"
	"omega/cmd/restapi/server"
	"omega/cmd/restapi/startoff"
	"omega/internal/core"
	"omega/internal/corstartoff"
	"omega/pkg/dict"
	"omega/pkg/glog"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {

	engine := startoff.LoadEnvs()

	glog.Init(engine.Envs[core.ServerLogFormat],
		engine.Envs[core.ServerLogOutput],
		engine.Envs[core.ServerLogLevel],
		engine.Envs.ToBool(core.ServerLogJSONIndent),
		true)

	dict.Init(engine.Envs[core.TermsPath], engine.Envs.ToBool(core.TranslateInBackend))

	corstartoff.ConnectDB(engine, false)
	corstartoff.ConnectActivityDB(engine)
	startoff.Migrate(engine)

	insertdata.Insert(engine)

	corstartoff.LoadSetting(engine)

	server.Start(engine)

}
