package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"omega/internal/core"
	"omega/server"
)

func main() {

	engine := core.StartEngine()
	defer engine.DB.Close()
	defer engine.ActivityDB.Close()

	s := server.Initialize(engine)

	engine.ServerLog.Info("Starting Server...")
	err := s.Run(fmt.Sprintf("%v:%v", engine.Environments.Server.ADDR, engine.Environments.Server.Port))
	if err != nil {
		engine.ServerLog.Fatal(err)
	}
	fmt.Printf("••• Server started \n\n")
	engine.ServerLog.Info("Server started.")
}
