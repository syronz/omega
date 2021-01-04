package main

import (
	"flag"
	"omega/cmd/restapi/insertdata"
	"omega/cmd/restapi/server"
	"omega/cmd/restapi/startoff"
	"omega/cmd/restapi/startoff/event"
	"omega/cmd/restapi/startoff/procedure"
	"omega/cmd/restapi/startoff/view"
	"omega/internal/core"
	"omega/internal/corstartoff"
	"omega/pkg/dict"
	"omega/pkg/glog"
	// _ "gorm.io/gorm/dialects/mysql"
	// _ "gorm.io/gorm/dialects/postgres"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

//var memprofile = flag.String("memprofile", "../performance/mem.prof", "write memory profile to `file`")

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

	//init of views
	view.InitViewReports(engine)
	view.InitDasboardViews(engine)

	//init of procedures
	procedure.InitDashboardProcedure(engine)
	procedure.InitReportProcedure(engine)
	//init of events
	event.InitdashboardEvent(engine)
	event.InitreportEvent(engine)

	corstartoff.LoadSetting(engine)

	server.Start(engine)

}
