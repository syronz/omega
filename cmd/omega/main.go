package main

import (
	"fmt"
	"omega/initiate"
	"omega/server"
)

func main() {

	cfg, env := initiate.Setup()
	defer cfg.DB.Close()

	s := server.Setup(cfg)

	err := s.Run(fmt.Sprintf("%v:%v", env.Server.ADDR, env.Server.Port))
	if err != nil {
		cfg.Log.Fatal(err)
	}
}
