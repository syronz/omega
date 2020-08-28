package main

import (
	"omega/cmd/restapi/insertdata"
	"omega/cmd/restapi/server"
	"omega/cmd/restapi/startoff"
	"omega/internal/core"
	"omega/internal/initiate"
	"omega/pkg/dict"
	"omega/pkg/glog"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	engine := startoff.LoadEnvs()

	glog.Init(engine.Envs[core.ServerLogFormat],
		engine.Envs[core.ServerLogOutput],
		engine.Envs[core.ServerLogLevel],
		engine.Envs.ToBool(core.ServerLogJSONIndent),
		true)
	glog.Debug("hello")

	// log2 := glog.New(engine.Envs[core.APILogFormat],
	// 	engine.Envs[core.APILogOutput],
	// 	engine.Envs[core.APILogLevel],
	// 	engine.Envs.ToBool(core.APILogJSONIndent),
	// 	true)

	dict.Init(engine.Envs[core.TermsPath])

	// logparam.APILog(engine)
	// initiate.LoadTerms(engine)
	initiate.ConnectDB(engine, false)
	initiate.ConnectActivityDB(engine)

	startoff.Migrate(engine)
	insertdata.Insert(engine)

	initiate.LoadSetting(engine)

	server.Start(engine)

}
