package main

import (
	"fmt"
	"omega/initiate"
	"omega/server"
)

func main() {

	engine := initiate.Setup()
	defer engine.DB.Close()

	s := server.Setup(engine)

	err := s.Run(fmt.Sprintf("%v:%v", engine.Environments.Server.ADDR, engine.Environments.Server.Port))
	if err != nil {
		engine.Log.Fatal(err)
	}
}
