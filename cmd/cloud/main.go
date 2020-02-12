package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"omega/internal/core"
	// "omega/server"
	"omega/cloud"
)

func main() {

	engine := core.StartEngine()
	defer engine.DB.Close()
	defer engine.ActivityDB.Close()

	s := cloud.Initialize(engine)

	engine.ServerLog.Info("Starting Cloud...")
	err := s.Run(fmt.Sprintf("%v:%v", engine.Environments.Cloud.ADDR, engine.Environments.Cloud.Port))
	if err != nil {
		engine.ServerLog.Fatal(err)
	}
	fmt.Printf("••• Cloud started \n\n")
	engine.ServerLog.Info("Cloud started.")
}
