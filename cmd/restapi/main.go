package main

import (
	"omega/cmd/restapi/determine"
	"omega/cmd/restapi/insertdata"
	"omega/cmd/restapi/server"
	"omega/internal/initiate"
	"omega/internal/logparam"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {

	engine := determine.LoadEnvs()

	logparam.ServerLog(engine)
	logparam.APILog(engine)
	initiate.LoadTerms(engine)
	initiate.ConnectDB(engine, false)
	initiate.ConnectActivityDB(engine)

	determine.Migrate(engine)
	insertdata.Insert(engine)

	initiate.LoadSetting(engine)

	server.Start(engine)

}
